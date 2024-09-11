package poke_questions

import (
	poke_types "github.com/zedosoad1995/pokemon-wordle/constants/pokemon/types"
	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
)

type possibleQuestionsBody struct {
	condition func(pokemon.Pokemon) bool
	text      string
}

func getAllQuestions() map[string]possibleQuestionsBody {
	allQuestions := map[string]possibleQuestionsBody{
		HasOnlyOneType: {
			condition: pokemon.HasOnlyOneType,
			text:      "Only has one type",
		},
		HasTwoTypes: {
			condition: pokemon.HasOnlyOneType,
			text:      "Has 2 types",
		},
	}

	for _, pokeType := range poke_types.AllPokeTypes {
		label := HasType(pokeType)
		text := "Type " + pokeType

		q := possibleQuestionsBody{
			text:      text,
			condition: pokemon.HasType(pokeType),
		}

		allQuestions[label] = q
	}

	return allQuestions
}
