package routes

import (
	"net/http"

	"gorm.io/gorm"
)

func CreateRoutes(mux *http.ServeMux, db *gorm.DB) {
	mux.HandleFunc("GET /api/pokemons", getPokemonsHandler(db))

	// TODO: call by id, this is just a test
	mux.HandleFunc("GET /api/boards", getBoardHandler(db))
}
