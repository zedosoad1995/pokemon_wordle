package pokemon

type PokemonList []Pokemon

func (pokemons PokemonList) Filter(conds ...func(Pokemon) bool) PokemonList {
	var res []Pokemon
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
