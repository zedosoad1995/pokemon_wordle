package pokemon

import (
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	PokedexNum  uint16
	Name        string
	Type1       *string
	Type2       *string
	Height      float64
	Weight      float64
	IsLegendary bool
	Gen         uint8
	BaseTotal   uint16
}

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
