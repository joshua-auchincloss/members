package server

import (
	"context"
	"crypto/tls"
	"members/common"
	"members/config"
	"net"

	"github.com/rs/zerolog/log"
)

type (
	Client interface {
		Dial(ctx context.Context, addr string) (net.Conn, error)
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

func (b *clientbase) Dial(ctx context.Context, addr string) (net.Conn, error) {
	return b.dialer(ctx, addr)
}
