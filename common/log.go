package common

import "github.com/rs/zerolog"

type (
	Logs interface {
		BuildLogger(root *zerolog.Logger)
		GetLogger() *zerolog.Logger
	}
)
