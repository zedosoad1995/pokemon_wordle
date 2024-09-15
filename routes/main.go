package routes

import (
	"net/http"

	"github.com/zedosoad1995/pokemon-wordle/middlewares"
	"gorm.io/gorm"
)

func CreateRoutes(mux *http.ServeMux, db *gorm.DB) {
	errorHandler := middlewares.ErrorHandler

	mux.HandleFunc("GET /api/pokemons", errorHandler(getPokemonsHandler(db)))

	mux.HandleFunc("GET /api/boards/{boardNum}", errorHandler(getBoardHandler(db)))

	mux.HandleFunc("PUT /api/boards/{boardNum}/answers", errorHandler(updateAnswerHandler(db)))
}
