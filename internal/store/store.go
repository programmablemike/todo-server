// a simple in-memory task store
package store

type Store[T any] struct {
	storage []T
}
type Comparator[T any] func(this T) bool

func NewStore[T any](v ...T) *Store[T] {
	if len(v) > 0 {
		return &Store[T]{
			storage: v,
		}
	} else {
		return &Store[T]{}
	}
}

// Adds a new value to the store
func (s *Store[T]) Add(v T) {
	s.storage = append(s.storage, v)
}

// Adds multiple values to the store
func (s *Store[T]) AddAll(v []T) {
	s.storage = append(s.storage, v...)
}

// DeleteOne deletes the first matching value returning true for found one and false if it didn't
func (s *Store[T]) DeleteOne(cmp Comparator[T]) bool {
	for idx, t := range s.storage {
		if cmp(t) {
			s.storage = append(s.storage[:idx], s.storage[idx+1:]...)
			return true
		}
	}
	return false
}

// DeleteAll deletes all matching values and returns the number of deleted items
func (s *Store[T]) DeleteAll(cmp Comparator[T]) int {
	count := 0
	for idx, t := range s.storage {
		if cmp(t) {
			s.storage = append(s.storage[:idx], s.storage[idx+1:]...)
			count++
		}
	}
	return count
}

// FindOne finds the first value to match the comparator and returns it
func (s *Store[T]) FindOne(cmp Comparator[T]) T {
	var result T
	for _, t := range s.storage {
		if cmp(t) {
			result = t
		}
	}
	return result
}

// FindAll finds all values matching the comparator and returns them as a slice
func (s *Store[T]) FindAll(cmp Comparator[T]) []T {
	var results []T
	for _, t := range s.storage {
		if cmp(t) {
			results = append(results, t)
		}
	}
	return results
}

// ListAll items in the store without any filtering
// Used for list calls and testing
func (s *Store[T]) ListAll() []T {
	return s.storage
}
