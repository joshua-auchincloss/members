package storage_fx

import (
	"members/storage"

	"go.uber.org/fx"
)

var (
	Dependencies = fx.Module(
		"storage",
		fx.Provide(New, storage.WithStore),
		fx.Invoke(storage.Setup),
	)

	Setup = fx.Module(
		"setup",
		fx.Invoke(with_create),
		Dependencies,
	)
)