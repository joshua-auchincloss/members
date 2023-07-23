package storage_fx

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/uptrace/bun"
)

type qh struct {
	sub *zerolog.Logger
}

func (qh *qh) BeforeQuery(c context.Context, qe *bun.QueryEvent) context.Context {
	return c
}

func (qh *qh) AfterQuery(c context.Context, qe *bun.QueryEvent) {
	ev := qh.sub.Debug().Str("sql", qe.Query)
	if qe.Err != nil {
		ev.Err(qe.Err).Bool("ok", false).Send()
	} else {
		ev.Bool("ok", true).Send()
	}
}
