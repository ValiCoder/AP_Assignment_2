package store

import "sync"

// Generic repository for Part E requirement
type GenericRepository[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func NewGenericRepository[K comparable, V any]() *GenericRepository[K, V] {
	return &GenericRepository[K, V]{
		data: make(map[K]V),
	}
}

func (r *GenericRepository[K, V]) Create(key K, value V) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[key] = value
}

func (r *GenericRepository[K, V]) Get(key K) (V, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	value, exists := r.data[key]
	return value, exists
}

func (r *GenericRepository[K, V]) GetAll() []V {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]V, 0, len(r.data))
	for _, value := range r.data {
		result = append(result, value)
	}
	return result
}
