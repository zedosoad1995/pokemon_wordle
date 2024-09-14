package answer

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	BoardID    uint
	UserID     uint
	Cell11     *string
	Cell12     *string
	Cell13     *string
	Cell21     *string
	Cell22     *string
	Cell23     *string
	Cell31     *string
	Cell32     *string
	Cell33     *string
	TriesLeft  uint8
	IsGameOver bool
}
