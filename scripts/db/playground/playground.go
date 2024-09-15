package main

import (
	"fmt"

	db_pkg "github.com/zedosoad1995/pokemon-wordle/config/db"
	"github.com/zedosoad1995/pokemon-wordle/config/env"
	"github.com/zedosoad1995/pokemon-wordle/models/answer"
)

func main() {
	env.LoadEnvs()
	db := db_pkg.Init()

	err := answer.UpsertAnswer(db, 1, 1, "lala", "1", 1)
	fmt.Print(err)
}
