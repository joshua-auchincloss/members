package common

import (
	"context"
	"fmt"
	"members/utils"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/feature"
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

var (
	_ bun.AfterCreateTableHook = ((*Membership)(nil))
	_ bun.BeforeInsertHook     = ((*Membership)(nil))
)

func (m *Membership) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().
		Model((*Membership)(nil)).
		Unique().
		Index("membership_addr_port").
		Column("address", "service").Exec(ctx)
	return err
}

func (m *Membership) BeforeInsert(ctx context.Context, query *bun.InsertQuery) error {
	db := query.DB()
	if db.HasFeature(feature.InsertOnConflict) {
		*query = *query.On(
			"conflict (address, service) do update set last_health = excluded.last_health",
		)
	} else if db.HasFeature(feature.InsertOnDuplicateKey) {
		*query = *query.On(
			fmt.Sprintf(
				"duplicate key update %s = %s",
				utils.QuoteCol(db, "last_health"),
				utils.QuoteCol(db, "last_health"),
			),
		)
	}
	return nil
}
