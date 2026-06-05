package optional

// Optional represents a value that may or may not be set.
type Optional[T any] struct {
	value T
	isSet bool
}

// New creates a new Optional holding the provided value.
func New[T any](v T) Optional[T] {
	return Optional[T]{value: v, isSet: true}
}

// IsSet returns true if the value has been explicitly set.
func (o Optional[T]) IsSet() bool {
	return o.isSet
}

// Get returns the underlying value.
// If the Optional has not been set, the zero value of T is returned without
// any error or panic. Always call IsSet() before calling Get() to avoid
// acting on a meaningless zero value.
func (o Optional[T]) Get() T {
	return o.value
}
