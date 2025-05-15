package result

import "fmt"

// Result is a monadic type that represents a value or an error. It is used to
// handle errors in a functional way, allowing for chaining operations without
// the need for explicit error handling.
type Result[T any] struct {
	v   *T
	err error
}

// Value returns a Result with a value.
func Value[T any](v T) Result[T] { return Result[T]{v: &v} }

// Error returns a Result with an error.
func Error[T any](err error) Result[T] { return Result[T]{err: err} }

// IsError checks if the Result contains a error.
func IsError[T any](o Result[T]) bool { return o.err != nil }

// Must is a helper function that extracts the value from a Result. If the
// Result contains an error, it returns a zero value of the type and an error.
// This is useful for cases where you want to assert that the Result is
// successful and handle the error in a different way.
func Must[T any](r Result[T]) (T, error) {
	if IsError(r) {
		return *new(T), fmt.Errorf("Error[%T]: %w", *new(T), r.err)
	}

	return *r.v, nil
}
