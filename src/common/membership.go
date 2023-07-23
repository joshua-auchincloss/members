package common

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type (
	Membership struct {
		bun.BaseModel `bun:"members"`

		PublicAddress  string    `bun:"address"`
		Service        Service   `bun:"service"`
		JoinTime       time.Time `bun:"time,default:current_timestamp"`
		LastHealthTime time.Time `bun:"last_health,default:current_timestamp"`
	}
)

func (m *Membership) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().
		Model((*Membership)(nil)).
		Unique().
		Index("membership_addr_port").
		Column("address", "service").Exec(ctx)
	return err
}
