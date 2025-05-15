package result

// PureFunc is a function that takes an input of type TIn and returns an output
// of type TOut.
type PureFunc[TIn, TOut any] func(TIn) TOut

// Map applies a pure function to the value inside the Result if it is not an
// error. If the Result is an error, it returns the same error without applying
// the function.
func Map[TIn, TOut any](
	r Result[TIn],
	fn func(TIn) TOut,
) Result[TOut] {
	if IsError(r) {
		return Error[TOut](r.err)
	}

	return Value(fn(*r.v))
}
