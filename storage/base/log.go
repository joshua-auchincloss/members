package base

import (
	"members/config"
	"members/logging"

	"github.com/rs/zerolog"
)

func LoggerFor(prov config.ConfigProvider, root *zerolog.Logger) *zerolog.Logger {
	store := prov.GetConfig().Storage
	sub := logging.WithSub(root, "storage", func(ctx zerolog.Context) zerolog.Context {
		ctx = ctx.Str("storage-kind", store.Kind).
			Str("storage-uri", store.URI).
			Bool("storage-debug", store.Debug)
		if store.Debug {
			ctx = ctx.
				Uint32("storage-port", store.Port).
				Bool("storage-create-on-init", store.Create).
				Bool("storage-drop-on-init", store.Drop)
		}
		return ctx
	})
	return sub
}
