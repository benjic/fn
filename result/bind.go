package result

func Bind[TIn, TOut any](
	o Result[TIn],
	fn MonadicFunc[TIn, TOut],
) Result[TOut] {
	if IsError(o) {
		return Error[TOut](o.err)
	}

	return fn(*o.v)
}
