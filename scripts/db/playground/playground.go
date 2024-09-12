package main

import (
	db_pkg "github.com/zedosoad1995/pokemon-wordle/config/db"
	"github.com/zedosoad1995/pokemon-wordle/config/env"
	"github.com/zedosoad1995/pokemon-wordle/models/board"
)

func main() {
	env.LoadEnvs()
	db := db_pkg.Init()

	board.GetAnswers(db, 1)
}
