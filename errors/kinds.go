package errs

import "errors"

var (
	ErrorOccupied = errors.New("occupied by member")
)
