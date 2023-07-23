package utils

func OneOf[T any](candidates []T, ismatch func(T) bool) bool {
	for _, candidate := range candidates {
		if ismatch(candidate) {
			return true
		}
	}
	return false
}

func AnyEq[T comparable](candidates []T, check T) bool {
	return OneOf(
		candidates,
		func(in T) bool {
			return in == check
		},
	)
}
