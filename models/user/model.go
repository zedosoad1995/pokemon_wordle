package user

import (
	"github.com/zedosoad1995/pokemon-wordle/models/answer"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Token   string `gorm:"unique;not null"`
	Answers []answer.Answer
}
