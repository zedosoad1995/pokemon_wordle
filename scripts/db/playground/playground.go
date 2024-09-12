package main

import (
	db_pkg "github.com/zedosoad1995/pokemon-wordle/config/db"
	"github.com/zedosoad1995/pokemon-wordle/config/env"
	poke_questions "github.com/zedosoad1995/pokemon-wordle/constants/pokemon/questions"
	poke_types "github.com/zedosoad1995/pokemon-wordle/constants/pokemon/types"
	"github.com/zedosoad1995/pokemon-wordle/models/board"
)

func main() {
	env.LoadEnvs()
	db := db_pkg.Init()

	newBoardBody := board.InsertBody{
		Col1: poke_questions.HasOnlyOneType,
		Col2: poke_questions.HasTwoTypes,
		Col3: poke_questions.HasType(poke_types.Flying),
		Row1: poke_questions.HasType(poke_types.Electric),
		Row2: poke_questions.HasType(poke_types.Fire),
		Row3: poke_questions.HasType(poke_types.Ice),
	}
	board.Insert(db, newBoardBody)
}
