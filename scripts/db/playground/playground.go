package main

import (
	"fmt"

	db_pkg "github.com/zedosoad1995/pokemon-wordle/config/db"
	"github.com/zedosoad1995/pokemon-wordle/config/env"
	poke_types "github.com/zedosoad1995/pokemon-wordle/constants/pokemon/types"
	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
)

func main() {
	env.LoadEnvs()
	db := db_pkg.Init()

	pokemons := pokemon.GetPokemonsByGen(db, 1)
	pokemons = pokemons.Filter(pokemon.HasType(poke_types.Water), pokemon.HasTwoTypes)

	var pokemonsNames []string
	for _, p := range pokemons {
		pokemonsNames = append(pokemonsNames, p.Name)
	}

	fmt.Print(pokemonsNames, len(pokemonsNames))
}
