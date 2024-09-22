package routes

import (
	"errors"
	"net/http"
	"strconv"

	poke_questions "github.com/zedosoad1995/pokemon-wordle/constants/pokemon/questions"
	"github.com/zedosoad1995/pokemon-wordle/models/answer"
	"github.com/zedosoad1995/pokemon-wordle/models/board"
	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
	"github.com/zedosoad1995/pokemon-wordle/models/user"
	route_types "github.com/zedosoad1995/pokemon-wordle/routes/types"
	"github.com/zedosoad1995/pokemon-wordle/utils"
	"gorm.io/gorm"
)

type getAnswersFreqRes struct {
	Freqs [3][3]map[string]uint `json:"freqs"`
}

func getAnswersFreq(db *gorm.DB) route_types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		boardNumStr := r.PathValue("boardNum")
		boardNum, err := strconv.ParseUint(boardNumStr, 10, 0)
		if err != nil {
			return utils.SendJSON(w, 400, route_types.ErrorRes{
				Message: "invalid board number",
			})
		}

		query := `
			WITH answer_cells AS (
				SELECT 
					answers.cell11, 
					answers.cell12, 
					answers.cell13, 
					answers.cell21, 
					answers.cell22, 
					answers.cell23, 
					answers.cell31, 
					answers.cell32, 
					answers.cell33
				FROM answers
				JOIN boards
					ON boards.id = answers.board_id
				WHERE boards.board_num = ?
			)`

		cells := [9]string{"cell11", "cell12", "cell13", "cell21", "cell22", "cell23", "cell31", "cell32", "cell33"}
		for i, cell := range cells {
			query += `
				SELECT '` + cell + `' AS cell, ` + cell + ` AS pokemon, COUNT(` + cell + `) AS freq
				FROM answer_cells
				GROUP BY ` + cell

			if i < len(cells)-1 {
				query += `
				UNION ALL`
			}
		}

		var pokeFreqs []struct {
			Cell    string
			Pokemon string
			Freq    uint
		}
		if err := db.Raw(query, boardNum).Scan(&pokeFreqs).Error; err != nil {
			return err
		}

		var res [3][3]map[string]uint
		for i := range res {
			for j := range res[i] {
				res[i][j] = make(map[string]uint)
			}
		}

		cellToRow := map[string]int{"cell11": 1, "cell12": 1, "cell13": 1, "cell21": 2, "cell22": 2, "cell23": 2, "cell31": 3, "cell32": 3, "cell33": 3}
		cellToCol := map[string]int{"cell11": 1, "cell12": 2, "cell13": 3, "cell21": 1, "cell22": 2, "cell23": 3, "cell31": 1, "cell32": 2, "cell33": 3}
		for _, freq := range pokeFreqs {
			if freq.Pokemon == "" {
				continue
			}

			rowNum := cellToRow[freq.Cell] - 1
			colNum := cellToCol[freq.Cell] - 1

			res[rowNum][colNum][freq.Pokemon] = freq.Freq
		}

		return utils.SendJSON(w, 200, getAnswersFreqRes{
			Freqs: res,
		})
	}
}

type updateAnswerBody struct {
	UserToken string `json:"userToken"`
	Row       int    `json:"row"`
	Col       int    `json:"col"`
	Answer    string `json:"answer"`
}

func updateAnswerHandler(db *gorm.DB) route_types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		boardNumStr := r.PathValue("boardNum")
		// TODO: body validation
		body, err := utils.GetJSONBody[updateAnswerBody](r)
		if err != nil {
			return err
		}

		boardNum, err := strconv.ParseUint(boardNumStr, 10, 0)
		if err != nil {
			return utils.SendJSON(w, 400, route_types.ErrorRes{
				Message: "invalid board number",
			})
		}

		userRes, err := user.GetUserByToken(db, body.UserToken)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.SendJSON(w, 400, route_types.ErrorRes{
					Message: "invalid user_token",
				})
			}

			return err
		}

		var res []struct {
			BoardID    uint
			Row1       string
			Row2       string
			Row3       string
			Col1       string
			Col2       string
			Col3       string
			AnswerID   uint
			IsGameOver bool
		}
		query := `
			SELECT 
				boards.id AS board_id,
				boards.col1,
				boards.col2,
				boards.col3,
				boards.row1,
				boards.row2,
				boards.row3,
				answers.id AS answer_id,
				answers.is_game_over
			FROM boards
			LEFT JOIN answers 
				ON boards.id = answers.board_id AND answers.user_id = ?
			WHERE boards.board_num = ?
			LIMIT 1`
		if err := db.Raw(query, userRes.ID, boardNum).Scan(&res).Error; err != nil {
			return err
		}

		if len(res) == 0 {
			return utils.SendJSON(w, 400, route_types.ErrorRes{
				Message: "invalid board number",
			})
		}

		if res[0].IsGameOver {
			return utils.SendJSON(w, 400, route_types.ErrorRes{
				Message: "Answers have already been submitted.",
			})
		}

		// TODO: maybe have another DB with all the answers, so it is more performant to check the answer
		pokemons, err := pokemon.GetPokemons(db)
		if err != nil {
			return err
		}

		rows := []string{res[0].Row1, res[0].Row2, res[0].Row3}
		cols := []string{res[0].Col1, res[0].Col2, res[0].Col3}
		filteredPokemons := pokemons.Filter(
			poke_questions.AllQuestions[rows[body.Row-1]].Condition,
			poke_questions.AllQuestions[cols[body.Col-1]].Condition,
		)

		isAnswerValid := utils.Some(filteredPokemons, func(p pokemon.Pokemon) bool {
			return p.Name == body.Answer
		})
		if !isAnswerValid {
			return utils.SendJSON(w, 400, route_types.ErrorRes{
				Message: "invalid answer",
			})
		}

		if err := answer.UpsertSingleCell(db, body.Row, body.Col, body.Answer, userRes.ID, res[0].BoardID, res[0].AnswerID); err != nil {
			return err
		}

		return utils.SendJSON(w, 200, route_types.SuccessRes{
			Message: "Updated answer.",
		})
	}
}

type updateAnswersBody struct {
	UserToken string     `json:"userToken"`
	Answers   [][]string `json:"answers"`
}

func updateAnswersHandler(db *gorm.DB) route_types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		boardNumStr := r.PathValue("boardNum")
		// TODO: body validation
		body, err := utils.GetJSONBody[updateAnswersBody](r)
		if err != nil {
			return err
		}

		boardNum, err := strconv.ParseUint(boardNumStr, 10, 0)
		if err != nil {
			return utils.SendJSON(w, 400, route_types.ErrorRes{
				Message: "invalid board number",
			})
		}

		userRes, err := user.GetUserByToken(db, body.UserToken)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.SendJSON(w, 400, route_types.ErrorRes{
					Message: "invalid user_token",
				})
			}

			return err
		}

		boardRes, err := board.GetBoardByNum(db, uint(boardNum))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.SendJSON(w, 400, route_types.ErrorRes{
					Message: "invalid board number",
				})
			}
			return err
		}

		answerRes, err := answer.GetAnswer(db, userRes.ID, boardRes.ID)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			answerRes, err = answer.CreateAnswer(db, userRes.ID, boardRes.ID)
			if err != nil {
				return err
			}
		}

		if answerRes.IsGameOver {
			return utils.SendJSON(w, 400, route_types.ErrorRes{
				Message: "Answers have already been submitted.",
			})
		}

		pokemons, err := pokemon.GetPokemons(db)
		if err != nil {
			return err
		}

		validAnswers, err := board.GetValidAnswers(db, *boardRes, pokemons)
		if err != nil {
			return err
		}

		mapAnswers := [3][3]*string{
			{(*answerRes).Cell11,
				(*answerRes).Cell12,
				(*answerRes).Cell13},
			{(*answerRes).Cell21,
				(*answerRes).Cell22,
				(*answerRes).Cell23},
			{(*answerRes).Cell31,
				(*answerRes).Cell32,
				(*answerRes).Cell33},
		}

		answerSetters := [3][3]func(val *string){
			{func(val *string) { answerRes.Cell11 = val },
				func(val *string) { answerRes.Cell12 = val },
				func(val *string) { answerRes.Cell13 = val }},
			{func(val *string) { answerRes.Cell21 = val },
				func(val *string) { answerRes.Cell22 = val },
				func(val *string) { answerRes.Cell23 = val }},
			{func(val *string) { answerRes.Cell31 = val },
				func(val *string) { answerRes.Cell32 = val },
				func(val *string) { answerRes.Cell33 = val }},
		}

		for row := uint8(0); row < 3; row++ {
			for col := uint8(0); col < 3; col++ {
				// If not nil check if body is sending the same value (shouldn't be able to edit)
				if mapAnswers[row][col] != nil {
					if *mapAnswers[row][col] != body.Answers[row][col] {
						return utils.SendJSON(w, 400, route_types.ErrorRes{
							Message: "invalid answer",
						})
					}
					continue
				}

				isAnswerValid := body.Answers[row][col] == "" || utils.Some(validAnswers[row][col], func(p string) bool {
					return p == body.Answers[row][col]
				})

				if !isAnswerValid {
					return utils.SendJSON(w, 400, route_types.ErrorRes{
						Message: "invalid answer",
					})
				}

				if body.Answers[row][col] != "" {
					answerSetters[row][col](&body.Answers[row][col])
				}
			}
		}

		answerRes.IsGameOver = true
		if err := db.Save(&answerRes).Error; err != nil {
			return err
		}

		return utils.SendJSON(w, 200, route_types.SuccessRes{
			Message: "Updated answers.",
		})
	}
}
