package storage

type Cache[K comparable, V any] interface {
	Get(key K) (V, error)
	Set(key K, value V)
	Delete(key K)
}
