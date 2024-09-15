package routes

import (
	"errors"
	"fmt"
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

type BoardRes struct {
	Cols []string
	Rows []string
}

type GetBoardHandlerRes struct {
	Answers board.Answers
	Board   BoardRes
}

func getBoardHandler(db *gorm.DB) route_types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		boardNumStr := r.PathValue("boardNum")
		boardNum, err := strconv.ParseUint(boardNumStr, 10, 0)
		if err != nil {
			return utils.SendJSON(w, 400, route_types.ErrorRes{
				Message: "Invalid board number",
			})
		}

		boardObj, err := board.GetBoardByNum(db, uint(boardNum))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.SendJSON(w, 404, route_types.ErrorRes{
					Message: fmt.Sprintf("Board number %d does not exist", boardNum),
				})
			}

			return err
		}

		pokemons, err := pokemon.GetPokemonsByGen(db, 1)
		if err != nil {
			return err
		}

		answers, err := board.GetAnswers(db, *boardObj, pokemons)
		if err != nil {
			return err
		}

		cols := []string{boardObj.Col1, boardObj.Col2, boardObj.Col3}
		rows := []string{boardObj.Row1, boardObj.Row2, boardObj.Row3}

		transformedCols := utils.Map(cols, func(col string) string {
			return poke_questions.AllQuestions[col].Text
		})
		transformedRows := utils.Map(rows, func(row string) string {
			return poke_questions.AllQuestions[row].Text
		})

		boardRes := BoardRes{
			Cols: transformedCols,
			Rows: transformedRows,
		}

		return utils.SendJSON(w, 200, GetBoardHandlerRes{Answers: *answers, Board: boardRes})
	}
}

type updateAnswerBody struct {
	UserToken string `json:"user_token"`
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

		userId := strconv.FormatUint(uint64(userRes.ID), 10)

		var res []struct {
			BoardID  uint
			Row1     string
			Row2     string
			Row3     string
			Col1     string
			Col2     string
			Col3     string
			AnswerID uint
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
				answers.id AS answer_id
			FROM boards
			LEFT JOIN answers 
				ON boards.id = answers.board_id AND answers.user_id = ?
			WHERE boards.board_num = ?
			LIMIT 1`
		if err := db.Raw(query, userId, boardNum).Scan(&res).Error; err != nil {
			return err
		}

		if len(res) == 0 {
			return utils.SendJSON(w, 400, route_types.ErrorRes{
				Message: "invalid board number",
			})
		}

		// TODO: maybe have another DB with all the answers, so it is more performant to check the answer
		pokemons, err := pokemon.GetPokemonsByGen(db, 1)
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

		if err := answer.UpsertAnswer(db, body.Row, body.Col, body.Answer, userId, res[0].BoardID, res[0].AnswerID); err != nil {
			return err
		}

		return utils.SendJSON(w, 200, "")
	}
}
