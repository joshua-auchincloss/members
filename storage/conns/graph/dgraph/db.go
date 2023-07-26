package dgraph

import (
	"members/config"
	"net"
	"strconv"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	dgraph_port uint32 = 8080
	dgraph_host        = "localhost"
	dgraph_user        = "groot"
	dgraph_pass        = "password"
)

func New(prov config.ConfigProvider) (*dgo.Dgraph, error) {
	cfg := prov.GetConfig()
	cfg.Storage.OverrideIfNull(
		dgraph_user,
		dgraph_pass,
		dgraph_port,
		dgraph_host,
	)
	var tr credentials.TransportCredentials
	if cfg.Storage.SSL {
		tlscfg := cfg.Storage.Tls
		if tlscfg != nil {
			certs, err := tlscfg.Build()
			if err != nil {
				return nil, err
			}
			tr = credentials.NewTLS(
				certs,
			)
		}
	} else {
		tr = insecure.NewCredentials()
	}
	d, err := grpc.Dial(
		net.JoinHostPort(cfg.Storage.URI, strconv.Itoa(int(cfg.Storage.Port))),
		grpc.WithTransportCredentials(tr),
	)
	if err != nil {
		return nil, err
	}
	dc := dgo.NewDgraphClient(api.NewDgraphClient(d))
	// err = dc.Login(context.TODO(), user, pass)
	return dc, err
}
