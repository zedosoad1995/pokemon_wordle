package board

import (
	"time"

	"gorm.io/gorm"
)

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

func Insert(db *gorm.DB, body InsertBody) (bool, error) {
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
			return false, res.Error
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
		return false, err
	}

	return true, nil
}
