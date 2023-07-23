package app

import (
	"members/config"
	errs "members/errors"
	"members/logging"
	"members/p2p"
	"members/service"
	"members/service/health"
	"members/service/registry"
	"members/service/starter"
	storage_fx "members/storage/fx"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

var cmds = []*cli.Command{
	{
		Name:  "start",
		Flags: config.Flags(),
		Action: func(orig *cli.Context) error {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			app := fx.New(
				fx.Supply(orig, sigs),
				logging.Module,
				errs.Module,
				config.Module,
				storage_fx.Dependencies,
				service.Module,
				health.Module,
				registry.Module,
				starter.Module,
				p2p.Module,
			)
			if err := app.Start(orig.Context); err != nil {
				return (err)
			}
			<-sigs
			if err := app.Stop(orig.Context); err != nil {
				return (err)
			}
			return nil
		},
	},
	{
		Name:  "schema",
		Flags: config.Flags(),
		Action: func(ctx *cli.Context) error {
			app := fx.New(
				fx.Supply(ctx),
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
}

func closer() {
	os.Exit(0)
}

func BuildApp() *cli.App {
	log.Print("ok")
	app := cli.NewApp()
	app.Name = "mm"
	app.Usage = ""
	app.Version = "0.1"
	app.Suggest = true
	app.EnableBashCompletion = true
	app.Commands = cmds
	return app
}
