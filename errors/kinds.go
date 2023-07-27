package errs

import "errors"

var (
	ErrSurvive = errors.New("poison")
	ErrPoison  = errors.New("poison")

	ErrOccupied      = errors.Join(ErrPoison, errors.New("occupied by member"))
	ErrCastInvalid   = errors.Join(ErrPoison, errors.New("type cast"))
	ErrServerStarter = errors.Join(ErrPoison, errors.New("server starter"))
)
