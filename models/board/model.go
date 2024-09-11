package board

import (
	"time"

	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	BoardNum uint32
	Date     time.Time `gorm:"type:date"`
	Col1     string
	Col2     string
	Col3     string
	Row1     string
	Row2     string
	Row3     string
}
