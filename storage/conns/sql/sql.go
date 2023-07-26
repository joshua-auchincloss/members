package sql

import (
	"context"
	"database/sql"
	"fmt"
	"members/common"
	"members/config"
	"members/storage/base"
	"members/utils"
	"time"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

type (
	SqlBase struct {
		DB     *bun.DB
		logger *zerolog.Logger
		kind   common.StorageType
		prov   config.ConfigProvider
	}
	SqlInitializer = func(prov config.ConfigProvider) (*bun.DB, error)
	SqlConverter   = func(*bun.DB) (base.BaseStore, error)
)

var (
	_ base.BaseStore = ((*SqlBase)(nil))
)

func create_sql(
	prov config.ConfigProvider,
	factory base.StoreFactory,
	root *zerolog.Logger,
) (base.BaseStore, error) {
	sub := base.LoggerFor(prov, root)
	if db, err := factory.New(); err != nil {
		return nil, err
	} else {
		db.WithLogger(sub)
		if prov.GetConfig().Storage.Debug {

			// db..AddQueryHook(&qh{sub})
		}
		return db, nil
	}
}

func NewSqlBase(db *bun.DB, log *zerolog.Logger, typ common.StorageType) (base.BaseStore, error) {
	return &SqlBase{db, log, typ, nil}, nil
}

func (s *SqlBase) Kind() common.StorageType {
	return s.kind
}

func (s *SqlBase) WithProvider(prov config.ConfigProvider) {
	s.prov = prov
}

func (s *SqlBase) WithQueryHook(i interface{}) {
	s.DB.AddQueryHook(i.(bun.QueryHook))
}

func (s *SqlBase) Logger() *zerolog.Logger {
	return s.logger
}

func (sq *SqlBase) WithLogger(sub *zerolog.Logger) {
	sq.logger = sub
}

func runCreate[T interface{}](
	sq *SqlBase,
	model T,
	chain ...func(*bun.CreateTableQuery) *bun.CreateTableQuery,
) error {
	var fc func(*bun.CreateTableQuery) *bun.CreateTableQuery
	if len(chain) != 0 {
		fc = chain[0]
	}
	base := sq.DB.NewCreateTable().
		Model(model).
		IfNotExists()
	if fc != nil {
		base = fc(base)
	}
	if _, err := base.Exec(context.TODO()); err != nil {
		return err
	}
	return nil
}

func drop[T interface{}](sq *SqlBase, model T) {
	if _, err := sq.DB.
		NewDropTable().
		Model(model).
		Cascade().
		Exec(context.TODO()); err != nil {
		sq.Logger().Print(err)
	}
}

func (sq *SqlBase) QuoteCol(v string) string {
	return utils.QuoteCol(sq.DB, v)
}

func (sq *SqlBase) Template(str string, vs ...any) string {
	return fmt.Sprintf(str, vs...)
}

func (sq *SqlBase) CleanOldMembers(ctx context.Context, from time.Duration) error {
	if _, err := sq.DB.NewDelete().
		Model((*common.Membership)(nil)).
		Where("last_health < ?", time.Now().Add(-from)).
		Exec(ctx); err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (sq *SqlBase) UpsertMembership(ctx context.Context, meta *common.Membership) error {
	if meta.Dns == "" {
		meta.Dns = meta.PublicAddress
	}

	base := sq.DB.NewInsert().
		Model(meta)
	if _, err := base.
		Exec(context.TODO(), meta); err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (sq *SqlBase) GetMembers(ctx context.Context, kind ...common.Service) ([]*common.Membership, error) {
	ok := utils.OkSpread(kind)
	var out []*common.Membership
	base := sq.DB.NewSelect().
		Model(&out)
	if ok {
		base = base.Where(
			"service in (?)", bun.In(kind),
		)
	}
	if err := base.Scan(context.TODO(), &out); err != nil {
		sq.logger.Print(err)
		return nil, err
	}
	return out, nil
}

func (sq *SqlBase) Teardown(ctx context.Context) error {
	for _, tbl := range []interface{}{
		(*common.Membership)(nil),
		(*common.ProtoProject)(nil),
		(*common.RegisteredProto)(nil),
		(*common.ProtoMeta)(nil),
	} {
		drop(sq, tbl)
	}
	return nil
}

func (sq *SqlBase) create() error {
	if err := runCreate(sq, (*common.Membership)(nil)); err != nil {
		return err
	}
	if err := runCreate(sq, (*common.ProtoProject)(nil)); err != nil {
		return err
	}
	if err := runCreate(sq, (*common.ProtoMeta)(nil), func(ctq *bun.CreateTableQuery) *bun.CreateTableQuery {
		return ctq.ForeignKey(
			sq.Template(
				"(%s) references repos (%s) on update cascade on delete cascade",
				sq.QuoteCol("key"),
				sq.QuoteCol("id"),
			),
		)
	}); err != nil {
		return err
	}
	if err := runCreate(sq, (*common.RegisteredProto)(nil), func(ctq *bun.CreateTableQuery) *bun.CreateTableQuery {
		return ctq.ForeignKey(
			sq.Template(
				"(%s) references registration_meta (%s) on update cascade on delete cascade",
				sq.QuoteCol("id"),
				sq.QuoteCol("id"),
			))
	}); err != nil {
		return err
	}
	return nil
}

func (sq *SqlBase) Setup(ctx context.Context, cfg config.ConfigProvider) error {
	storg := cfg.GetConfig().Storage
	if storg.Drop {
		if err := sq.Teardown(ctx); err != nil {
			return err
		}
	}
	if storg.Create {
		if err := sq.create(); err != nil {
			return err
		}
	}
	return nil
}
