package main

import (
	db_pkg "github.com/zedosoad1995/pokemon-wordle/config/db"
	"github.com/zedosoad1995/pokemon-wordle/config/env"
)

func main() {
	env.LoadEnvs()
	db_pkg.Init()
}
