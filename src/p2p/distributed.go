package p2p

type (
	distributed[T any] struct {
		scaling_factor int
		workload       []byte
	}
)
