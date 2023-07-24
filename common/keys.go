package common

type (
	Key[K, V comparable] struct {
		m map[K]V
	}
)

func NewKey[K, V comparable](m map[K]V) Key[K, V] {
	return Key[K, V]{
		m,
	}
}
func (m *Key[K, V]) Get(k K) V {
	return m.m[k]
}

func (m *Key[K, V]) Reverse() map[V]K {
	o := make(map[V]K)
	for k, v := range m.m {
		o[v] = k
	}
	return o
}
