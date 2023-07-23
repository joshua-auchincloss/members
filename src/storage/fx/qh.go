package storage_fx

import (
	"context"
	"log"

	"github.com/uptrace/bun"
)

type qh struct {
}

func (qh *qh) BeforeQuery(c context.Context, qe *bun.QueryEvent) context.Context {
	log.Print(qe.Query)
	return c
}
func (qh *qh) AfterQuery(c context.Context, qe *bun.QueryEvent) {
}
