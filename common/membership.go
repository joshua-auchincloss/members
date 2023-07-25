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
		bun.BaseModel  `bun:"members"`
		Dns            string    `bun:"dns" json:"dns"`
		PublicAddress  string    `bun:"address" json:"address"`
		Service        Service   `bun:"service" json:"service"`
		JoinTime       time.Time `bun:"registration,default:current_timestamp" json:"registration"`
		LastHealthTime time.Time `bun:"last_health,default:current_timestamp" json:"last_health"`
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
