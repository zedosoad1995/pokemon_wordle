package answer

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func GetAnswer(db *gorm.DB, userId, boardId uint) (*Answer, error) {
	var answerRes Answer

	if err := db.Where(&Answer{UserID: userId, BoardID: boardId}).First(&answerRes).Error; err != nil {
		return nil, err
	}

	return &answerRes, nil
}

func GetAnswers(db *gorm.DB, boardId uint) ([]Answer, error) {
	var answersRes []Answer

	if err := db.Where(&Answer{BoardID: boardId}).Find(&answersRes).Error; err != nil {
		return nil, err
	}

	return answersRes, nil
}

func (currAnswer Answer) CalculateScore(freqs [3][3]map[string]uint) float64 {
	cells := [][]*string{
		{currAnswer.Cell11, currAnswer.Cell12, currAnswer.Cell13},
		{currAnswer.Cell21, currAnswer.Cell22, currAnswer.Cell23},
		{currAnswer.Cell31, currAnswer.Cell32, currAnswer.Cell33},
	}

	score := 0.0

	for i, row := range cells {
		for j, cell := range row {
			if cell == nil {
				score += 100
			} else {
				score += float64(freqs[i][j][*cell])
			}
		}
	}

	return score
}

func CountAnswersFromBoard(db *gorm.DB, boardId uint) (*uint, error) {
	var count int64

	if err := db.Model(&Answer{}).Where("board_id = ?", boardId).Count(&count).Error; err != nil {
		return nil, err
	}

	convCount := uint(count)
	return &convCount, nil
}

func CreateAnswer(db *gorm.DB, userId, boardId uint) (*Answer, error) {
	var createdAnswer Answer
	answerToCreate := Answer{BoardID: boardId, UserID: userId}

	if err := db.Create(answerToCreate).Scan(&createdAnswer).Error; err != nil {
		return nil, err
	}

	return &createdAnswer, nil
}

func UpsertSingleCell(db *gorm.DB, row, col int, answer string, userId, boardId, answerId uint) error {
	colName := fmt.Sprintf("cell%d%d", row, col)
	now := time.Now()

	if answerId == 0 {
		query := `
			INSERT INTO answers (board_id, user_id, ` + colName + `, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?)`
		if err := db.Exec(query, boardId, userId, answer, now, now).Error; err != nil {
			return err
		}

		return nil
	}

	query := `
		UPDATE answers
		SET ` + colName + ` = ?, updated_at = ?
		WHERE id = ?`
	err := db.Exec(query, answer, now, answerId).Error
	if err != nil {
		return err
	}

	return nil
}
