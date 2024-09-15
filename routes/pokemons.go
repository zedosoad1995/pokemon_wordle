package routes

import (
	"encoding/json"
	"net/http"

	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
	route_types "github.com/zedosoad1995/pokemon-wordle/routes/types"
	"github.com/zedosoad1995/pokemon-wordle/utils"
	"gorm.io/gorm"
)

type GetPokemonsRes struct {
	Pokemons []string
}

func getPokemonsHandler(db *gorm.DB) route_types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		pokemons, err := pokemon.GetPokemonsByGen(db, 1)
		if err != nil {
			return err
		}

		pokemonNames := utils.Map(pokemons, func(p pokemon.Pokemon) string {
			return p.Name
		})

		json.NewEncoder(w).Encode(GetPokemonsRes{Pokemons: pokemonNames})

		return nil
	}
}
