package pokemon

import "gorm.io/gorm"

func GetPokemonsByGen(db *gorm.DB, gen int) (PokemonList, error) {
	var res []Pokemon

	err := db.Raw(`
		SELECT *
		FROM pokemons
		WHERE gen = ?
	`, gen).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetPokemons(db *gorm.DB) (PokemonList, error) {
	return GetPokemonsByGen(db, 1);
}
