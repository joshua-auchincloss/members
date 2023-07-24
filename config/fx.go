package config

import (
	"go.uber.org/fx"
)

var (
	Module = fx.Module(
		"config",
		fx.Provide(
			New,
		),
	)

	CliCacheCfg = fx.Module(
		"cli-cache-cfg",
		Module,
		fx.Invoke(ensureFileCache("cli.cache")),
	)
)
