package board

import (
	"time"

	poke_questions "github.com/zedosoad1995/pokemon-wordle/constants/pokemon/questions"
	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
	"github.com/zedosoad1995/pokemon-wordle/utils"
	"gorm.io/gorm"
)

func GetBoardByNum(db *gorm.DB, boardNum uint) (*Board, error) {
	var board Board
	err := db.Where("board_num = ?", boardNum).First(&board).Error

	if err != nil {
		return nil, err
	}

	return &board, nil
}

type InsertBody struct {
	BoardNum *uint32
	Date     *time.Time
	Col1     string
	Col2     string
	Col3     string
	Row1     string
	Row2     string
	Row3     string
}

func Insert(db *gorm.DB, body InsertBody) error {
	// If no boardNum or date defined, it will go to the next boardNum/date
	if body.BoardNum == nil || body.Date == nil {
		type getBoard struct {
			BoardNum uint32
			Date     time.Time
		}
		var lastBoard getBoard
		res := db.Raw(`
			SELECT board_num, date
			FROM boards
			ORDER BY board_num DESC
			LIMIT 1
		`).Scan(&lastBoard)

		if res.Error != nil {
			return res.Error
		}

		if body.BoardNum == nil {
			boardNum := uint32(lastBoard.BoardNum) + 1
			body.BoardNum = &boardNum
		}

		if body.Date == nil {
			var nextDate time.Time
			if res.RowsAffected == 0 {
				nextDate = time.Now().Truncate(24 * time.Hour)
			} else {
				nextDate = lastBoard.Date.Add(24 * time.Hour)
			}
			body.Date = &nextDate
		}
	}

	newBoard := Board{
		BoardNum: *body.BoardNum,
		Date:     *body.Date,
		Col1:     body.Col1,
		Col2:     body.Col2,
		Col3:     body.Col3,
		Row1:     body.Row1,
		Row2:     body.Row2,
		Row3:     body.Row3,
	}
	if err := db.Create(&newBoard).Error; err != nil {
		return err
	}

	return nil
}

type Answers [3][3][]string

func GetValidAnswers(db *gorm.DB, board Board, pokemons pokemon.PokemonList) (*Answers, error) {
	rows := []string{board.Row1, board.Row2, board.Row3}
	cols := []string{board.Col1, board.Col2, board.Col3}

	var answers Answers
	for i, row := range rows {
		for j, col := range cols {
			rowLabel, rowVal := utils.ExtractLabelAndValue(row)
			colLabel, colVal := utils.ExtractLabelAndValue(col)

			filteredPokemons := pokemons.Filter(
				poke_questions.AllQuestions[rowLabel](rowVal).Condition,
				poke_questions.AllQuestions[colLabel](colVal).Condition,
			)

			answers[i][j] = utils.Map(filteredPokemons, func(p pokemon.Pokemon) string {
				return p.Name
			})
		}
	}

	return &answers, nil
}
