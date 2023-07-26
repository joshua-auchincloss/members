package common

import "members/common/internal"

type (
	StorageType = int
)

const (
	StorageMemory StorageType = iota + internal.StorageOffset
	StoragePostgres
	StorageMysql
	StorageDGraph
	StorageNeo4j
)
