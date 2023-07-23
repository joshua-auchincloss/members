package logging

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var (
	Module = fx.Module(
		"logging",
		fx.Provide(new_writer, new_root),
	)
)

func new_writer() io.Writer {
	return os.Stdout
}

func new_root(read io.Writer) *zerolog.Logger {
	log := zerolog.New(read)
	return &log
}

func WithSub(root *zerolog.Logger, service string, apply ...func(zerolog.Context) zerolog.Context) *zerolog.Logger {
	var fu func(zerolog.Context) zerolog.Context
	if len(apply) != 0 {
		fu = apply[0]
	}
	w := root.With().Str("service", service)
	if fu != nil {
		w = fu(w)
	}
	log := w.Logger()
	return &log
}
