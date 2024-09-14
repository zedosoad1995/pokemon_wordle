package routes

import (
	"net/http"

	"gorm.io/gorm"
)

func CreateRoutes(mux *http.ServeMux, db *gorm.DB) {
	mux.HandleFunc("GET /api/pokemons", getPokemonsHandler(db))
}
