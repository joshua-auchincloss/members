package app

import (
	"context"
	"errors"
	"fmt"
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
	"members/utils"
	"os"
	"os/signal"
	"syscall"

	stdlog "log"

	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

var (
	head = color.New(color.FgMagenta).Add(color.Underline).Add(color.Bold)
)

func nooplog(ctx *cli.Context, opts ...fx.Option) fx.Option {
	out := []fx.Option{}
	if !ctx.Bool("debug") {
		opts = append(opts, fx.NopLogger)
	}
	out = append(out, opts...)
	return fx.Options(out...)
}

func resolver_printer(prov config.ConfigProvider) {
	fmt.Println(head.Sprint("\n\ncluster from dynamic configuration\n"))
	tbl := utils.Table("DNS", "Service", "Address")
	for svc, clust := range prov.GetDynamic().PeekAll() {
		for dns, addrs := range clust.Peek() {
			for _, addr := range addrs {
				tbl.AddRow(dns, common.ServiceKeys.Get(svc), addr.Addr)
			}
		}

	}
	tbl.Print()
}

func remoteInvoker[V any](
	clientfactory fx.Option,
	do func(config.ConfigProvider, *zerolog.Logger, V) error,
) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		dial := &common.DialArgs{
			DNS:       ctx.String(config.AdminCliDnsDefn.Key),
			Addresses: ctx.StringSlice(config.AdminCliAddrDefn.Key),
			TLS:       ctx.Bool(config.RemoteTlsDefn.Key),
		}
		opts := nooplog(ctx,
			fx.Supply(ctx, dial),
			logging.Fx,
			logging.Suppress,
			fx.Invoke(
				func(da *common.DialArgs) {
					if ctx.Bool(config.RemoteDebugDefn.Key) {
						log.Debug().Interface("dial", dial).Send()
					}
				},
			),
			config.CliCacheCfg,
			storage_fx.Dependencies,
			clientfactory,
			fx.Invoke(resolver_printer),
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
					if err := do(prov, root, c); err != nil {
						return err
					}
					return nil
				},
			),
			fx.Invoke(closer),
		)
		app := fx.New(
			opts,
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
				Flags: config.RemoteFlags(),
				Action: remoteInvoker[health.HealthClient](
					health_impl.ClientFactory, func(prov config.ConfigProvider, root *zerolog.Logger, c health.HealthClient) error {
						resp, err := c.Check(context.TODO(), &healthpb.HealthCheckRequest{
							Service: "",
						})
						root.Debug().Err(err).Interface("resp", resp).Send()
						if err != nil {
							return err
						}
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
				Flags: config.RemoteFlags(),
				Action: remoteInvoker[admin.AdminClient](
					admin_impl.ClientFactory,
					func(prov config.ConfigProvider, root *zerolog.Logger, c admin.AdminClient) error {
						resp, err := c.DescribeCluster(context.TODO(), &commonpb.Empty{})
						if err != nil {
							return err
						}

						fmt.Println(head.Sprint("\n\ncluster response\n"))
						tbl := utils.Table("DNS", "Service", "Address", "Join Time", "Last Health")
						for {
							membr, err := resp.Recv()
							eof := errors.Is(err, io.EOF)
							if err != nil && !eof {
								return err
							} else if eof {
								break
							}
							tbl.AddRow(membr.Dns, common.ServiceKeys.Get(membr.Service), membr.Address, membr.JoinTime, membr.LastHealth)
							root.Debug().Err(err).Interface("member", membr).Send()
						}
						tbl.Print()
						return nil
					},
				),
			},
			{
				Name:  "start",
				Flags: config.ClusterFlags(),
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
