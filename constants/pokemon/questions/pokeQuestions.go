package poke_questions

import (
	"github.com/zedosoad1995/pokemon-wordle/utils"
)

const (
	HasOnlyOneType = "HasOnlyOneType"
	HasTwoTypes    = "HasTwoTypes"
)

func HasType(pokeType string) string {
	return "Has" + utils.CapitalizeFirstLetter(pokeType) + "Type"
}

var AllQuestions = getAllQuestions()

var QuestionLabels = utils.GetKeys(AllQuestions)
