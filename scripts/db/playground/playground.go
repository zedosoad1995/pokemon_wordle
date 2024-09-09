package main

import (
	"fmt"

	db_pkg "github.com/zedosoad1995/pokemon-wordle/config/db"
	"github.com/zedosoad1995/pokemon-wordle/config/env"
	"github.com/zedosoad1995/pokemon-wordle/constants"
	"gorm.io/gorm"
)

type PokemonList []db_pkg.Pokemon

func getPokemonsByGen(db *gorm.DB, gen int) PokemonList {
	var res []db_pkg.Pokemon

	db.Raw(`
		SELECT *
		FROM pokemons
		WHERE gen = ?
	`, gen).Scan(&res)

	return res
}

func hasOnlyOneType(p db_pkg.Pokemon) bool {
	return p.Type1 != nil && p.Type2 == nil
}

func hasTwoTypes(p db_pkg.Pokemon) bool {
	return p.Type1 != nil && p.Type2 != nil
}

func isTypeWater(p db_pkg.Pokemon) bool {
	return (p.Type1 != nil && *p.Type1 == constants.Water) || (p.Type2 != nil && *p.Type2 == constants.Water)
}

func (pokemons PokemonList) filterPokemons(conds ...func(db_pkg.Pokemon) bool) PokemonList {
	var res []db_pkg.Pokemon
	for _, p := range pokemons {
		isValid := true

		for _, cond := range conds {
			if !cond(p) {
				isValid = false
				break
			}
		}

		if isValid {
			res = append(res, p)
		}
	}
	return res
}

func main() {
	env.LoadEnvs()
	db := db_pkg.Init()

	pokemons := getPokemonsByGen(db, 1)
	pokemons = pokemons.filterPokemons(isTypeWater, hasTwoTypes)

	var pokemonsNames []string
	for _, p := range pokemons {
		pokemonsNames = append(pokemonsNames, p.Name)
	}

	fmt.Print(pokemonsNames, len(pokemonsNames))
}
