package poke_questions

import (
	"github.com/zedosoad1995/pokemon-wordle/utils"
)

const (
	HasOnlyOneType      = "HasOnlyOneType"
	HasTwoTypes         = "HasTwoTypes"
	HasType             = "HasType"
	StartsWithLetter    = "StartsWithLetter"
	NameHasLenGreaterEq = "NameHasLenGreaterEq"
	NameHasLenLessEq    = "NameHasLenLessEq"
	HeightGreaterEq     = "HeightGreaterEq"
	HeightLessEq        = "HeightLessEq"
	WeightGreaterEq     = "WeightGreaterEq"
	WeightLessEq        = "WeightLessEq"
)

var AllQuestions = getAllQuestions()

var QuestionLabels = utils.GetKeys(AllQuestions)
