package utils

import (
	"context"
	"time"

	"go.uber.org/fx"
)

func FxWithTermination(app *fx.App, ctx context.Context) error {
	sigs := Sigs()
	if err := app.Start(ctx); err != nil {
		return (err)
	}
	try_close := func() error {
		ctx, clean := context.WithTimeout(ctx, time.Second)
		defer clean()
		if err := app.Stop(ctx); err != nil {
			return err
		}
		return nil
	}
	return WaitForTermination(try_close, sigs)
}
