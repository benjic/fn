package optional

import "fmt"

var (
	ErrMissingValue = fmt.Errorf("optional value required")
)

type Optional[T any] struct {
	value *T
}

func Value[T any](value T) Optional[T] { return Optional[T]{value: &value} }

func Empty[T any]() Optional[T] { return Optional[T]{value: nil} }

func HasValue[T any](o Optional[T]) bool { return o.value != nil }

func IsEmpty[T any](o Optional[T]) bool { return o.value == nil }

func Must[T any](o Optional[T]) (T, error) {
	if IsEmpty(o) {
		return *new(T), fmt.Errorf("Optional[%T]: %w", *new(T), ErrMissingValue)
	}

	return *o.value, nil
}
