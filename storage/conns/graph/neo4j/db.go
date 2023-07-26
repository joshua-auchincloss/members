package neo4j

import (
	"members/config"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	neo4j_user        = "neo4j"
	neo4j_pass        = "neo4j"
	neo4j_port uint32 = 7687 // bolt
	neo4j_host        = "localhost"
)

func New(prov config.ConfigProvider) (neo4j.DriverWithContext, error) {
	cfg := prov.GetConfig()
	cfg.Storage.OverrideIfNull(
		neo4j_user,
		neo4j_pass,
		neo4j_port,
		neo4j_host,
	)
	at := neo4j.BasicAuth(cfg.Storage.Username, cfg.Storage.Password, "")
	db, err := neo4j.NewDriverWithContext(cfg.Storage.URI, at)
	if err != nil {
		return nil, err
	}
	return db, nil
}
