package optional

func Bind[TIn, TOut any](
	o Optional[TIn],
	fn MonadicFunc[TIn, TOut],
) Optional[TOut] {
	if HasValue(o) {
		return fn(*o.value)
	}

	return Empty[TOut]()
}
