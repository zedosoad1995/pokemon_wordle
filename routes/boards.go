package routes

import (
	"encoding/json"
	"net/http"

	poke_questions "github.com/zedosoad1995/pokemon-wordle/constants/pokemon/questions"
	"github.com/zedosoad1995/pokemon-wordle/models/board"
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

func getBoardHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		boardObj, _ := board.GetBoardByNum(db, 1)
		if boardObj == nil {
			// TODO: not found
		}

		answers, _ := board.GetAnswers(db, *boardObj)

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

		json.NewEncoder(w).Encode(GetBoardHandlerRes{Answers: *answers, Board: boardRes})
	}
}
