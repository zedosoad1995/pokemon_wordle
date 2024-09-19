package routes

import (
	"net/http"

	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
	route_types "github.com/zedosoad1995/pokemon-wordle/routes/types"
	"github.com/zedosoad1995/pokemon-wordle/utils"
	"gorm.io/gorm"
)

type GetPokemonsRes struct {
	Pokemons []string `json:"pokemons"`
}

func getPokemonsHandler(db *gorm.DB) route_types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		pokemons, err := pokemon.GetPokemons(db)
		if err != nil {
			return err
		}

		pokemonNames := utils.Map(pokemons, func(p pokemon.Pokemon) string {
			return p.Name
		})

		return utils.SendJSON(w, 200, GetPokemonsRes{Pokemons: pokemonNames})
	}
}
