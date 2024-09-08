package main

import (
	"github.com/zedosoad1995/pokemon-wordle/config/db"
	"github.com/zedosoad1995/pokemon-wordle/config/env"
)

func main() {
	env.LoadEnvs()
	db.Init()
}
