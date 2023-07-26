package sql

import (
	"members/common"
	"members/storage/base"

	"github.com/uptrace/bun"
)

func SqlConvert(
	typ common.StorageType,
	conv ...func(*bun.DB) (base.BaseStore, error),
) SqlConverter {
	var do func(*bun.DB) (base.BaseStore, error)
	if len(conv) != 0 {
		do = conv[0]
	} else {
		do = base.NoOpBase(typ, NewSqlBase)
	}
	return func(d *bun.DB) (base.BaseStore, error) {
		if db, err := do(d); err != nil {
			return nil, err
		} else {
			return db, nil
		}
	}

}
