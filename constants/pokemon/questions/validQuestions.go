package poke_questions

import poke_types "github.com/zedosoad1995/pokemon-wordle/constants/pokemon/types"

/*
const (
	NameHasLenLessEq    = "NameHasLenLessEq"
)
*/

var ValidQuestions = []string{
	AllQuestions[HasOnlyOneType]("").Label,
	AllQuestions[HasTwoTypes]("").Label,
	AllQuestions[HasType](poke_types.Bug).Label,
	AllQuestions[HasType](poke_types.Dark).Label,
	AllQuestions[HasType](poke_types.Dragon).Label,
	AllQuestions[HasType](poke_types.Electric).Label,
	AllQuestions[HasType](poke_types.Fairy).Label,
	AllQuestions[HasType](poke_types.Fighting).Label,
	AllQuestions[HasType](poke_types.Fire).Label,
	AllQuestions[HasType](poke_types.Flying).Label,
	AllQuestions[HasType](poke_types.Ghost).Label,
	AllQuestions[HasType](poke_types.Grass).Label,
	AllQuestions[HasType](poke_types.Ground).Label,
	AllQuestions[HasType](poke_types.Ice).Label,
	AllQuestions[HasType](poke_types.Normal).Label,
	AllQuestions[HasType](poke_types.Poison).Label,
	AllQuestions[HasType](poke_types.Psychic).Label,
	AllQuestions[HasType](poke_types.Rock).Label,
	AllQuestions[HasType](poke_types.Steel).Label,
	AllQuestions[HasType](poke_types.Water).Label,
	AllQuestions[HeightGreaterEq]("1").Label,
	AllQuestions[HeightGreaterEq]("1.5").Label,
	AllQuestions[HeightGreaterEq]("2").Label,
	AllQuestions[HeightLessEq]("0.5").Label,
	AllQuestions[HeightLessEq]("1").Label,
	AllQuestions[WeightGreaterEq]("50").Label,
	AllQuestions[WeightGreaterEq]("80").Label,
	AllQuestions[WeightGreaterEq]("100").Label,
	AllQuestions[WeightLessEq]("5").Label,
	AllQuestions[WeightLessEq]("10").Label,
	AllQuestions[WeightLessEq]("20").Label,
	AllQuestions[StartsWithLetter]("M").Label,
	AllQuestions[StartsWithLetter]("P").Label,
	AllQuestions[StartsWithLetter]("S").Label,
	AllQuestions[StartsWithLetter]("G").Label,
	AllQuestions[StartsWithLetter]("D").Label,
	AllQuestions[NameHasLenEq]("7").Label,
	AllQuestions[NameHasLenEq]("6").Label,
	AllQuestions[NameHasLenEq]("8").Label,
	AllQuestions[NameHasLenEq]("9").Label,
	AllQuestions[NameHasLenGreaterEq]("8").Label,
	AllQuestions[NameHasLenGreaterEq]("9").Label,
	AllQuestions[NameHasLenLessEq]("5").Label,
	AllQuestions[NameHasLenLessEq]("6").Label,
}
