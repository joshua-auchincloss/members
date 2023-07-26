package dgraph

import (
	"members/common"
	"members/storage/base"

	"github.com/dgraph-io/dgo"
)

func GraphConvert(
	typ common.StorageType,
	conv ...func(*dgo.Dgraph) (base.BaseStore, error),
) GraphConverter {
	var do func(*dgo.Dgraph) (base.BaseStore, error)
	if len(conv) != 0 {
		do = conv[0]
	} else {
		do = base.NoOpBase(typ, NewDgraphBase)
	}
	return func(d *dgo.Dgraph) (base.BaseStore, error) {
		if db, err := do(d); err != nil {
			return nil, err
		} else {
			return db, nil
		}
	}

}
