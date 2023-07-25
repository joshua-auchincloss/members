package server

import (
	"context"
	"crypto/tls"
	"members/config"
	"time"
)

type (
	base struct {
		err  chan error
		srv  invariant
		tls  *tls.Config
		cfg  *config.ServerTls
		root *config.TlsConfig
	}
)

var (
	_ Server = ((*base)(nil))
)

func (s *base) GetServer() invariant {
	return s.srv
}

func (s *base) GetConfig() *config.ServerTls {
	return s.cfg
}

func (s *base) TlsEnabled() bool {
	return s.root.Enabled
}

func (s *base) GetTLS() *tls.Config {
	return s.tls
}

func (s *base) ReportErr(err error) {
	s.err <- err
}

func (s *base) Start() chan error {
	return starter(s)
}

func (s *base) Stop(ctx context.Context, ttc time.Duration) error {
	return stopper(s, ctx, ttc)
}
