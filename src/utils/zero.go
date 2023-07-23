package utils

func IsZero[T comparable](v T) bool {
	nw := new(T)
	return *nw == v
}

func ZeroStr(v string) bool {
	return v == ""
}
