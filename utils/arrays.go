package utils

func Includes[T comparable](s []T, elem T) bool {
	for _, item := range s {
		if item == elem {
			return true
		}
	}
	return false
}
