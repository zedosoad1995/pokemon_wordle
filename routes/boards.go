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
	route_types "github.com/zedosoad1995/pokemon-wordle/routes/types"
	"github.com/zedosoad1995/pokemon-wordle/utils"
	"gorm.io/gorm"
)

type BoardRes struct {
	Cols []string `json:"cols"`
	Rows []string `json:"rows"`
}

type GetBoardHandlerRes struct {
	Answers    board.Answers `json:"answers"`
	Board      BoardRes      `json:"board"`
	TotalPlays uint          `json:"totalPlays"`
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

		pokemons, err := pokemon.GetPokemons(db)
		if err != nil {
			return err
		}

		answers, err := board.GetValidAnswers(db, *boardObj, pokemons)
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

		count, err := answer.CountAnswersFromBoard(db, boardObj.ID)
		if err != nil {
			return err
		}

		boardRes := BoardRes{
			Cols: transformedCols,
			Rows: transformedRows,
		}

		return utils.SendJSON(w, 200, GetBoardHandlerRes{
			Answers:    *answers,
			Board:      boardRes,
			TotalPlays: *count,
		})
	}
}
