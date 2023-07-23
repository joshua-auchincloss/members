package cli

import (
	"log"
	"members/config"
	errs "members/errors"
	"members/service"
	"members/service/health"
	"members/service/registry"
	"members/service/starter"
	storage_fx "members/storage/fx"
	"os/signal"
	"syscall"

	// "members/raft"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

var cmds = []*cli.Command{
	{
		Name:  "start",
		Flags: config.Flags(),
		Action: func(ctx *cli.Context) error {
			app := fx.New(
				fx.Supply(ctx),
				errs.Module,
				fx.Provide(
					config.New,
				),
				storage_fx.Dependencies,
				service.Module,
				health.Module,
				registry.Module,
				starter.Module,
				// p2p.Module,
			)
			if err := app.Start(ctx.Context); err != nil {
				panic(err)
			}
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			<-sigs
			if err := app.Stop(ctx.Context); err != nil {
				panic(err)
			}
			return nil
		},
	},
	{
		Name:  "schema",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			app := fx.New(
				fx.Supply(ctx),
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
