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
