package poke_questions

import (
	poke_types "github.com/zedosoad1995/pokemon-wordle/constants/pokemon/types"
	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
)

type possibleQuestionsBody struct {
	Condition func(pokemon.Pokemon) bool
	Text      string
}

func getAllQuestions() map[string]possibleQuestionsBody {
	allQuestions := map[string]possibleQuestionsBody{
		HasOnlyOneType: {
			Condition: pokemon.HasOnlyOneType,
			Text:      "Only has one type",
		},
		HasTwoTypes: {
			Condition: pokemon.HasTwoTypes,
			Text:      "Has 2 types",
		},
	}

	for _, pokeType := range poke_types.AllPokeTypes {
		label := HasType(pokeType)
		text := "Type " + pokeType

		q := possibleQuestionsBody{
			Text:      text,
			Condition: pokemon.HasType(pokeType),
		}

		allQuestions[label] = q
	}

	return allQuestions
}
