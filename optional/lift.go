package optional

type MonadicFunc[A, B any] func(A) Optional[B]

// Lift allow you to promote any pure function into a monadic function.
func Lift[A, B any](fn PureFunc[A, B]) MonadicFunc[A, B] {
	return func(opt A) Optional[B] { return Value(fn(opt)) }
}
