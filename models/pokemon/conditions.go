package pokemon

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
