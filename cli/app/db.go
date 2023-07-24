package app

import (
	"members/config"
	"members/logging"
	storage_fx "members/storage/fx"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

var (
	database = cli.Command{
		Name: "database",
		Subcommands: []*cli.Command{
			{
				Name:  "schema",
				Flags: config.ClusterFlags(),
				Action: func(orig *cli.Context) error {
					app := fx.New(
						fx.Supply(orig),
						logging.Module,
						fx.Provide(
							config.New,
						),
						storage_fx.Setup,
						fx.Invoke(closer),
					)
					app.Run()
					return nil
				},
			},
		},
	}
)
