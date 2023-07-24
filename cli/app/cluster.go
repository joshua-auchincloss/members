package app

import (
	"context"
	"members/common"
	"members/config"
	errs "members/errors"
	"members/grpc/api/v1/health/pkg"
	"members/logging"
	"members/p2p"
	"members/service"
	"members/service/health"
	"members/service/registry"
	"members/service/starter"
	storage_fx "members/storage/fx"
	"os"
	"os/signal"
	"syscall"

	stdlog "log"

	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

var (
	cluster = cli.Command{
		Name: "cluster",
		Subcommands: []*cli.Command{
			{
				Name: "health",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "addr",
						Aliases: []string{"address", "a"},
						Value:   "localhost:9009",
					},
					&cli.BoolFlag{
						Name:    "tls",
						Aliases: []string{"enable-tls", "t"},
					},
				},
				Action: func(ctx *cli.Context) error {
					app := fx.New(
						fx.Supply(ctx, &common.DialArgs{
							Address: ctx.String("addr"),
							TLS:     ctx.Bool("tls"),
						}),
						logging.Module,
						config.CliCacheCfg,
						storage_fx.Dependencies,
						health.ClientFactory,
						fx.Invoke(
							func(
								root *zerolog.Logger,
								cf *health.HealthClient,
								prov config.ConfigProvider,
								args *common.DialArgs) error {
								cli, err := cf.New(prov, args)
								if err != nil {
									return err
								}
								c := *cli
								resp, err := c.Check(context.TODO(), &pkg.HealthCheckRequest{
									Service: "",
								})
								root.Debug().Err(err).Interface("resp", resp).Send()
								st := color.BlueString("status: ")
								switch {
								case resp.Status == pkg.HealthCheckResponse_SERVING:
									stdlog.Print(
										st, color.GreenString("serving"),
									)
								default:
									stdlog.Print(
										st, color.RedString("not serving"),
									)
								}
								return nil
							},
						),
						fx.Invoke(closer),
					)
					app.Run()
					return nil
				},
			},
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
		},
	}
)
