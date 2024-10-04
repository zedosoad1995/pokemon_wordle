package poke_questions

import (
	"strconv"

	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
)

type questionsRes struct {
	Condition func(pokemon.Pokemon) bool
	Text      string
	Label     string
}

type possibleQuestionsBody func(string) questionsRes

func getAllQuestions() map[string]possibleQuestionsBody {
	allQuestions := map[string]possibleQuestionsBody{
		HasOnlyOneType: func(value string) questionsRes {
			return questionsRes{
				Condition: pokemon.HasOnlyOneType,
				Text:      "Only has one type",
				Label:     "HasOnlyOneType",
			}
		},
		HasTwoTypes: func(value string) questionsRes {
			return questionsRes{
				Condition: pokemon.HasTwoTypes,
				Text:      "Has 2 types",
				Label:     "HasTwoTypes",
			}
		},
		HasType: func(value string) questionsRes {
			return questionsRes{
				Condition: pokemon.HasType(value),
				Text:      "Type " + value,
				Label:     "HasType:" + value,
			}
		},
		StartsWithLetter: func(value string) questionsRes {
			char := rune(value[0])

			return questionsRes{
				Condition: pokemon.NameStartsWithLetter(char),
				Text:      "Starts with the letter " + string(char),
				Label:     "StartsWithLetter:" + string(char),
			}
		},
		NameHasLenEq: func(value string) questionsRes {
			len, _ := strconv.Atoi(value)

			return questionsRes{
				Condition: pokemon.NameHasLenGreaterEq(len),
				Text:      "Name has " + value + " characters",
				Label:     "NameHasLenEq:" + value,
			}
		},
		NameHasLenGreaterEq: func(value string) questionsRes {
			len, _ := strconv.Atoi(value)

			return questionsRes{
				Condition: pokemon.NameHasLenGreaterEq(len),
				Text:      "Name has " + value + " characters or more",
				Label:     "NameHasLenGreaterEq:" + value,
			}
		},
		NameHasLenLessEq: func(value string) questionsRes {
			len, _ := strconv.Atoi(value)

			return questionsRes{
				Condition: pokemon.NameHasLenLessEq(len),
				Text:      "Name has " + value + " characters or less",
				Label:     "NameHasLenLessEq:" + value,
			}
		},
		HeightGreaterEq: func(value string) questionsRes {
			height, _ := strconv.ParseFloat(value, 64)

			// TODO: show conversion into feet

			return questionsRes{
				Condition: pokemon.HeightGreaterEq(height),
				Text:      "Has a height of " + value + "m or taller",
				Label:     "HeightGreaterEq:" + value,
			}
		},
		HeightLessEq: func(value string) questionsRes {
			height, _ := strconv.ParseFloat(value, 64)

			return questionsRes{
				Condition: pokemon.HeightLessEq(height),
				Text:      "Has a height of " + value + "m or shorter",
				Label:     "HeightLessEq:" + value,
			}
		},
		WeightGreaterEq: func(value string) questionsRes {
			weight, _ := strconv.ParseFloat(value, 64)

			// TODO: show conversion into kgs

			return questionsRes{
				Condition: pokemon.WeightGreaterEq(weight),
				Text:      "Has a weight of " + value + "kg or heavier",
				Label:     "WeightGreaterEq:" + value,
			}
		},
		WeightLessEq: func(value string) questionsRes {
			weight, _ := strconv.ParseFloat(value, 64)

			return questionsRes{
				Condition: pokemon.WeightLessEq(weight),
				Text:      "Has a weight of " + value + "kg or lighter",
				Label:     "WeightLessEq:" + value,
			}
		},
	}

	return allQuestions
}
