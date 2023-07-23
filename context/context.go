package wctx

import (
	"context"

	"go.uber.org/fx"
)

type (
	Context = context.Context

	ContextKeys = string
)

const (
	ContextKeyWithStore ContextKeys = "store"
)

func New() Context {
	return context.TODO()
}

var (
	NewContext = fx.Module("context",
		fx.Supply(
			fx.Annotate(New, fx.As(new(Context))),
		),
	)
)
