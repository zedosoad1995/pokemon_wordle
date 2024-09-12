package board

import (
	"time"

	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	BoardNum uint32    `gorm:"unique;not null"`
	Date     time.Time `gorm:"type:date;unique"`
	Col1     string
	Col2     string
	Col3     string
	Row1     string
	Row2     string
	Row3     string
}
