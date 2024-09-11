package pokemon

import "github.com/zedosoad1995/pokemon-wordle/constants"

type possibleQuestionsBody struct {
	condition func(Pokemon) bool
	text      string
}

var PossibleQuestions = map[string]possibleQuestionsBody{
	constants.HasOnlyOneType: {
		condition: Pokemon.HasOnlyOneType,
		text:      "Only has one type",
	},
	constants.HasTwoTypes: {
		condition: Pokemon.HasOnlyOneType,
		text:      "Has 2 types",
	},
	constants.HasWaterType: {
		condition: Pokemon.HasOnlyOneType,
		text:      "Type water",
	},
}

func (p Pokemon) HasOnlyOneType() bool {
	return p.Type1 != nil && p.Type2 == nil
}

func (p Pokemon) HasTwoTypes() bool {
	return p.Type1 != nil && p.Type2 != nil
}

func (p Pokemon) HasWaterType() bool {
	return (p.Type1 != nil && *p.Type1 == constants.Water) || (p.Type2 != nil && *p.Type2 == constants.Water)
}
