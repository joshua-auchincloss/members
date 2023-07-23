package common

type (
	Service = int
)

const (
	UnknownService = iota
	ServiceHealth
	ServiceRegistry
)
