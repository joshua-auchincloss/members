package mem

import (
	"database/sql"
	"fmt"
	"log"
	"members/config"
	"members/utils"

	"github.com/uptrace/bun"

	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

var (
	memtempl = "file::memory:?cache=shared"
	fitempl  = "file::%s?cache=shared&mode=memory"
)

func New(prov config.ConfigProvider) (*bun.DB, error) {
	fp := prov.GetConfig().Storage.URI

	log.Print("in mem")
	if utils.ZeroStr(fp) {
		fp = memtempl
	} else {
		fp = fmt.Sprintf(fitempl, fp)
	}
	sqldb, err := sql.Open(sqliteshim.ShimName, fp)
	if err != nil {
		return nil, err
	}
	return bun.NewDB(sqldb, sqlitedialect.New()), nil
}
