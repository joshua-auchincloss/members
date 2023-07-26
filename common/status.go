package common

import "members/common/internal"

type (
	Status = int
)

const (
	NoStatus      = iota
	StatusStarted = iota + internal.StatusOffset
	StatusClosed
)
