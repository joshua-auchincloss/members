package server

import (
	"context"
	"crypto/tls"
	"members/config"
	"time"
)

type (
	base struct {
		err     chan error
		srv     invariant
		tls     *tls.Config
		cfg     *config.ServerTls
		root    *config.TlsConfig
		starter func(Server) chan error

		closers []func() error
	}
)

func NewBase(
	err chan error,
	srv invariant,
	tls *tls.Config,
	cfg *config.ServerTls,
	root *config.TlsConfig,
	starter func(Server) chan error,
	closers ...func() error,
) base {
	return base{
		err,
		srv,
		tls,
		cfg,
		root,
		starter,
		closers,
	}
}

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
	return s.starter(s)
}

func (s *base) Stop(ctx context.Context, ttc time.Duration) error {
	return stopper(s, ctx, ttc)
}

func (s *base) WithCloser(f func() error) {
	s.closers = append(s.closers, f)
}
func (s *base) Closers() []func() error {
	return s.closers
}
