package utils

func OkSpread[T any](v []T) bool {
	return len(v) > 0
}

func Spread[T any](v []T, default_v ...T) T {
	var vo T
	if len(v) > 0 {
		vo = v[0]
	} else {
		if len(default_v) > 0 {
			vo = default_v[0]
		}
	}
	return vo
}
