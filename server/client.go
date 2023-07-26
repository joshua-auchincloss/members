package server

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"members/common"
	"members/config"
	"members/utils"
	"net"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type (
	Client interface {
		Dial(ctx context.Context, addr string) (net.Conn, error)
		DialFallback(ctx context.Context, addrs ...string) (net.Conn, error)
	}
	dialer = func(ctx context.Context, addr string, timeout time.Duration) (*net.Conn, error)

	clientbase struct {
		args *common.DialArgs
		tls  *tls.Config
		dialer
		svc common.Service
	}

	tcpClient struct {
		clientbase
	}
	udpClient struct {
		clientbase
	}
)

func NewClient(
	service common.Service,
	prov config.ConfigProvider,
	tlscfg *tls.Config,
	args *common.DialArgs,
) (Client, error) {
	cf := prov.GetConfig()
	log.Debug().Str("protocol", cf.Members.Protocol).Send()

	switch cf.Members.Protocol {
	case "udp":
		return &udpClient{clientbase{
			args,
			tlscfg,
			DialUDP,
			service,
		}}, nil
	default:
		log.Debug().Str("protocol", "tcp").Send()
		return &tcpClient{clientbase{
			args,
			tlscfg,
			DialTcp,
			service,
		}}, nil
	}
}

func baseDial(proto string) func(ctx context.Context, addr string, timeout time.Duration) (*net.Conn, error) {
	return func(ctx context.Context, addr string, timeout time.Duration) (*net.Conn, error) {
		return utils.LoopOrCancel(ctx, timeout, time.Millisecond, func() (net.Conn, error) {
			return net.Dial(proto, addr)
		})
	}
}

func DialTcp(ctx context.Context, addr string, timeout time.Duration) (*net.Conn, error) {
	return baseDial("tcp")(ctx, addr, timeout)
}

func DialUDP(ctx context.Context, addr string, timeout time.Duration) (*net.Conn, error) {
	return baseDial("udp")(ctx, addr, timeout)
}

func (cli *clientbase) Dial(ctx context.Context, addr string) (net.Conn, error) {
	if conn, err := cli.dialer(ctx, addr, time.Millisecond*1); err != nil {
		return nil, err
	} else {
		return *conn, nil
	}
}

func (cli *clientbase) DialFallback(ctx context.Context, addrs ...string) (net.Conn, error) {
	for _, addr := range addrs {
		this := addr
		if conn, err := utils.LoopOrCancel(ctx, time.Millisecond*100, time.Millisecond*1, func() (net.Conn, error) {
			return cli.Dial(ctx, this)
		}); conn != nil {
			return *conn, nil
		} else {
			log.Error().Err(err).Send()
		}
	}
	return nil, errors.New("no hosts available")
}

func DialGrpc(ctx context.Context, svc common.Service, dns string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts,
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	)
	return grpc.Dial(fmt.Sprintf("%s:///%s", common.ServiceKeys.Get(svc), dns),
		opts...)
}
