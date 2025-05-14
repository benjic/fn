package optional

type PureFunc[A, B any] func(A) B

func Map[A, B any](opt Optional[A], f PureFunc[A, B]) Optional[B] {
	if HasValue(opt) {
		return Value(f(*opt.value))
	}

	return Empty[B]()
}
