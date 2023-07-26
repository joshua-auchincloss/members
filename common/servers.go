package common

import "members/common/internal"

type (
	ServerKinds = int
)

const (
	ServerKindTCP = iota + internal.StorageOffset
	ServerKindUDP
)
