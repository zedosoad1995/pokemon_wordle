package utils

type mapFunc[E any, R any] func(E) R

func Map[S ~[]E, E any, R any](s S, f mapFunc[E, R]) []R {
	result := make([]R, len(s))
	for i := range s {
		result[i] = f(s[i])
	}
	return result
}
