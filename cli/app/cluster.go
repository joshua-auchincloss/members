package app

import (
	"context"
	"errors"
	"io"
	"members/common"
	"members/config"
	errs "members/errors"
	"members/grpc/api/v1/admin"
	commonpb "members/grpc/api/v1/common"
	"members/grpc/api/v1/health"
	healthpb "members/grpc/api/v1/health/pkg"
	"members/logging"
	"members/p2p"
	"members/service"
	admin_impl "members/service/admin"
	health_impl "members/service/health"
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
	remoteFlags = []cli.Flag{
		&cli.StringFlag{
			Name:    "addr",
			Aliases: []string{"address", "a"},
			Value:   "localhost:9009",
		},
		&cli.BoolFlag{
			Name:    "tls",
			Aliases: []string{"enable-tls", "t"},
		},
	}
)

func remoteInvoker[V any](
	clientfactory fx.Option,
	do func(*zerolog.Logger, V) error,
) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		app := fx.New(
			fx.Supply(ctx, &common.DialArgs{
				Address: ctx.String("addr"),
				TLS:     ctx.Bool("tls"),
			}),
			logging.Module,
			config.CliCacheCfg,
			storage_fx.Dependencies,
			clientfactory,
			fx.Invoke(
				func(
					root *zerolog.Logger,
					cf *service.ClientFactory[V],
					prov config.ConfigProvider,
					args *common.DialArgs) error {
					cli, err := cf.New(prov, args)
					if err != nil {
						return err
					}
					c := *cli
					if err := do(root, c); err != nil {
						return err
					}
					return nil
				},
			),
			fx.Invoke(closer),
		)
		app.Run()
		return nil
	}
}

var (
	cluster = cli.Command{
		Name: "cluster",
		Subcommands: []*cli.Command{
			{
				Name:  "health",
				Flags: remoteFlags,
				Action: remoteInvoker[health.HealthClient](
					health_impl.ClientFactory, func(root *zerolog.Logger, c health.HealthClient) error {
						resp, err := c.Check(context.TODO(), &healthpb.HealthCheckRequest{
							Service: "",
						})
						root.Debug().Err(err).Interface("resp", resp).Send()
						st := color.BlueString("status: ")
						switch {
						case resp.Status == healthpb.HealthCheckResponse_SERVING:
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
			},
			{
				Name:  "describe",
				Flags: remoteFlags,
				Action: remoteInvoker[admin.AdminClient](
					admin_impl.ClientFactory, func(root *zerolog.Logger, c admin.AdminClient) error {
						resp, err := c.DescribeCluster(context.TODO(), &commonpb.Empty{})
						if err != nil {
							return err
						}
						for {
							membr, err := resp.Recv()
							eof := errors.Is(err, io.EOF)
							if err != nil && !eof {
								return err
							} else if eof {
								break
							}
							root.Debug().Err(err).Interface("member", membr).Send()
						}
						return nil
					},
				),
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
						health_impl.Module,
						admin_impl.Module,
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
