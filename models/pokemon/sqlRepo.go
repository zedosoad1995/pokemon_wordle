package pokemon

import "gorm.io/gorm"

func GetPokemonsByGen(db *gorm.DB, gen int) PokemonList {
	var res []Pokemon

	db.Raw(`
		SELECT *
		FROM pokemons
		WHERE gen = ?
	`, gen).Scan(&res)

	return res
}
