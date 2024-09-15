package answer

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func UpsertAnswer(db *gorm.DB, row, col int, answer, userId string, boardId, answerId uint) error {
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
