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
	ImageUrl	*string
}
