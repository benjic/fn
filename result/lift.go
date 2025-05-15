package result

// MonadicFunc is a function that takes an input of type TIn and returns a
// Result of type TOut. It is used to represent a function that can be applied
// to a value inside a Result, allowing for chaining operations in a monadic
// style.
type MonadicFunc[TIn, TOut any] func(TIn) Result[TOut]

// Lift allow you to promote any pure function into a monadic function.
func Lift[TIn, TOut any](fn PureFunc[TIn, TOut]) MonadicFunc[TIn, TOut] {
	return func(opt TIn) Result[TOut] { return Value(fn(opt)) }
}
