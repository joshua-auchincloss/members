package logging

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func zapl(ctx *cli.Context) (*zap.Logger, error) {
	var log *zap.Logger
	var err error
	debug := ctx.Bool("debug")
	if debug {
		log = zap.NewNop()
	} else {
		log, err = zap.NewProduction()
	}
	return log, err
}

func new_lc(ctx *cli.Context) (fxevent.Logger, error) {
	var log *zap.Logger
	var err error
	debug := ctx.Bool("debug")
	if debug {
		log = zap.NewNop()
	} else {
		log, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	}
	zl := fxevent.ZapLogger{Logger: log}
	if !debug {
		zl.UseLogLevel(zapcore.DebugLevel)
	}
	return &zl, nil
}

var (
	Fx = fx.Module(
		"fx-suppression",
		fx.Provide(zapl),
		fx.WithLogger(new_lc),
	)

	Module = fx.Module(
		"logging",
		fx.Provide(new_writer, new_root),
	)

	Suppress = fx.Module(
		"logging-suppression",
		fx.Invoke(err_only), Module,
	)
)

func err_only(ctx *cli.Context) error {
	if !ctx.Bool("debug") {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
	return nil
}

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
