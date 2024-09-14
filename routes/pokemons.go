package routes

import (
	"encoding/json"
	"net/http"

	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
	"github.com/zedosoad1995/pokemon-wordle/utils"
	"gorm.io/gorm"
)

type GetPokemonsRes struct {
	Pokemons []string
}

func getPokemonsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		pokemons := pokemon.GetPokemonsByGen(db, 1)
		pokemonNames := utils.Map(pokemons, func(p pokemon.Pokemon) string {
			return p.Name
		})

		json.NewEncoder(w).Encode(GetPokemonsRes{Pokemons: pokemonNames})
	}
}
