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

var (
	lb_scheme = "round"
)

type (
	Client interface {
		Dial(ctx context.Context, addr string) (net.Conn, error)
		DialFallback(ctx context.Context, addrs ...string) (net.Conn, error)
	}
	dialer = func(ctx context.Context, addr string) (net.Conn, error)

	clientbase struct {
		args *common.DialArgs
		tls  *tls.Config
		dialer
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
	switch cf.Members.Protocol {
	case "udp":
		log.Debug().Str("protocol", "udp").Send()
		return &udpClient{clientbase{
			args,
			tlscfg,
			DialUDP,
		}}, nil
	default:
		log.Debug().Str("protocol", "tcp").Send()
		return &tcpClient{clientbase{
			args,
			tlscfg,
			DialTcp,
		}}, nil
	}
}

func DialTcp(ctx context.Context, addr string) (net.Conn, error) {
	return net.Dial("tcp", addr)
}

func DialUDP(ctx context.Context, addr string) (net.Conn, error) {
	return net.Dial("udp", addr)
}

func (cli *clientbase) Dial(ctx context.Context, addr string) (net.Conn, error) {
	return cli.dialer(ctx, addr)
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

func DialGrpc(ctx context.Context, dns string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts,
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	)
	return grpc.Dial(fmt.Sprintf("%s:///%s", lb_scheme, dns),
		opts...)
	// return nil, errors.New("no hosts available")
}
