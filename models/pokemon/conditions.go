package pokemon

import "unicode"

func HasOnlyOneType(p Pokemon) bool {
	return p.Type1 != nil && p.Type2 == nil
}

func HasTwoTypes(p Pokemon) bool {
	return p.Type1 != nil && p.Type2 != nil
}

func HasType(pokeType string) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return (p.Type1 != nil && *p.Type1 == pokeType) || (p.Type2 != nil && *p.Type2 == pokeType)
	}
}

func NameStartsWithLetter(char rune) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return unicode.ToLower(rune(p.Name[0])) == unicode.ToLower(char)
	}
}

func NameHasLenEq(length int) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return len(p.Name) == length
	}
}

func NameHasLenGreaterEq(length int) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return len(p.Name) >= length
	}
}

func NameHasLenLessEq(length int) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return len(p.Name) <= length
	}
}

func BaseTotalGreaterEq(min uint16) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.BaseTotal >= min
	}
}

func BaseTotalLessEq(max uint16) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.BaseTotal <= max
	}
}

func HeightGreaterEq(min float64) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.Height >= min
	}
}

func HeightLessEq(max float64) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.Height <= max
	}
}

func WeightGreaterEq(min float64) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.Weight >= min
	}
}

func WeightLessEq(max float64) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.Weight <= max
	}
}
