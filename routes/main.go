package routes

import (
	"net/http"

	"github.com/zedosoad1995/pokemon-wordle/middlewares"
	"gorm.io/gorm"
)

func CreateRoutes(mux *http.ServeMux, db *gorm.DB) {
	errorHandler := middlewares.ErrorHandler

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./public"))))

	mux.HandleFunc("GET /api/pokemons", errorHandler(getPokemonsHandler(db)))

	mux.HandleFunc("GET /api/boards/{boardNum}", errorHandler(getBoardHandler(db)))

	mux.HandleFunc("GET /api/boards/{boardNum}/answers/freq", errorHandler(getAnswersFreqHandler(db)))
	mux.HandleFunc("PUT /api/boards/{boardNum}/answers/one", errorHandler(updateAnswerHandler(db)))
	mux.HandleFunc("PUT /api/boards/{boardNum}/answers/submit", errorHandler(updateAnswersHandler(db)))

	mux.HandleFunc("POST /api/users", errorHandler(createUserHandler(db)))
}
