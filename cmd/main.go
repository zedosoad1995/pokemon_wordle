package main

import (
	"fmt"
	"net/http"

	db_pkg "github.com/zedosoad1995/pokemon-wordle/config/db"
	"github.com/zedosoad1995/pokemon-wordle/config/env"
	"github.com/zedosoad1995/pokemon-wordle/routes"
)

func main() {
	env.LoadEnvs()
	db := db_pkg.Init()

	mux := http.NewServeMux()
	routes.CreateRoutes(mux, db)

	var handler http.Handler = mux

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", handler)
}
