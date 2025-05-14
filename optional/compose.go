package optional

// Compose is a function that takes two monadic functions and returns a single
// monadic function that represents the composition of the two functions. It
// allows you to chain multiple monadic functions together in a functional
// style, where the output of one function is passed as the input to the next
// function. This is useful for creating complex transformations on values while
// maintaining the monadic structure of the Result type.
func Compose[A, B, C any](
	fAB MonadicFunc[A, B],
	fBC MonadicFunc[B, C],
) MonadicFunc[A, C] {
	return func(a A) Optional[C] {
		return Bind(Bind(Value(a), fAB), fBC)
	}

}

// Compose3 is a function that takes three monadic functions and returns a
// single monadic function that represents the composition of the three
// functions.
//
// This is a workaround for Go's type system which does not allow for dynamic
// function types in a list.
func Compose3[A, B, C, D any](
	fAB MonadicFunc[A, B],
	fBC MonadicFunc[B, C],
	fCD MonadicFunc[C, D],
) MonadicFunc[A, D] {
	return Compose(Compose(fAB, fBC), fCD)
}

// Compose4 is a function that takes four monadic functions and returns a single
// monadic function that represents the composition of the four functions.
//
// This is a workaround for Go's type system which does not allow for dynamic
// function types in a list.
func Compose4[A, B, C, D, E any](
	fAB MonadicFunc[A, B],
	fBC MonadicFunc[B, C],
	fCD MonadicFunc[C, D],
	fDE MonadicFunc[D, E],
) MonadicFunc[A, E] {
	return Compose(Compose3(fAB, fBC, fCD), fDE)
}

// Compose5 is a function that takes five monadic functions and returns a single
// monadic function that represents the composition of the five functions.
//
// This is a workaround for Go's type system which does not allow for dynamic
// function types in a list.
func Compose5[A, B, C, D, E, F any](
	fAB MonadicFunc[A, B],
	fBC MonadicFunc[B, C],
	fCD MonadicFunc[C, D],
	fDE MonadicFunc[D, E],
	fEF MonadicFunc[E, F],
) MonadicFunc[A, F] {
	return Compose(Compose4(fAB, fBC, fCD, fDE), fEF)
}

// Compose6 is a function that takes six monadic functions and returns a single
// monadic function that represents the composition of the six functions.
//
// This is a workaround for Go's type system which does not allow for dynamic
// function types in a list.
func Compose6[A, B, C, D, E, F, G any](
	fAB MonadicFunc[A, B],
	fBC MonadicFunc[B, C],
	fCD MonadicFunc[C, D],
	fDE MonadicFunc[D, E],
	fEF MonadicFunc[E, F],
	fFG MonadicFunc[F, G],
) MonadicFunc[A, G] {
	return Compose(Compose5(fAB, fBC, fCD, fDE, fEF), fFG)
}

// Compose7 is a function that takes seven monadic functions and returns a single
// monadic function that represents the composition of the seven functions.
//
// This is a workaround for Go's type system which does not allow for dynamic
// function types in a list.
func Compose7[A, B, C, D, E, F, G, H any](
	fAB MonadicFunc[A, B],
	fBC MonadicFunc[B, C],
	fCD MonadicFunc[C, D],
	fDE MonadicFunc[D, E],
	fEF MonadicFunc[E, F],
	fFG MonadicFunc[F, G],
	fGH MonadicFunc[G, H],
) MonadicFunc[A, H] {
	return Compose(Compose6(fAB, fBC, fCD, fDE, fEF, fFG), fGH)
}

// Compose8 is a function that takes eight monadic functions and returns a single
// monadic function that represents the composition of the eight functions.
//
// This is a workaround for Go's type system which does not allow for dynamic
// function types in a list.
func Compose8[A, B, C, D, E, F, G, H, I any](
	fAB MonadicFunc[A, B],
	fBC MonadicFunc[B, C],
	fCD MonadicFunc[C, D],
	fDE MonadicFunc[D, E],
	fEF MonadicFunc[E, F],
	fFG MonadicFunc[F, G],
	fGH MonadicFunc[G, H],
	fHI MonadicFunc[H, I],
) MonadicFunc[A, I] {
	return Compose(Compose7(fAB, fBC, fCD, fDE, fEF, fFG, fGH), fHI)
}

// Compose9 is a function that takes nine monadic functions and returns a single
// monadic function that represents the composition of the nine functions.
//
// This is a workaround for Go's type system which does not allow for dynamic
// function types in a list.
func Compose9[A, B, C, D, E, F, G, H, I, J any](
	fAB MonadicFunc[A, B],
	fBC MonadicFunc[B, C],
	fCD MonadicFunc[C, D],
	fDE MonadicFunc[D, E],
	fEF MonadicFunc[E, F],
	fFG MonadicFunc[F, G],
	fGH MonadicFunc[G, H],
	fHI MonadicFunc[H, I],
	fIJ MonadicFunc[I, J],
) MonadicFunc[A, J] {
	return Compose(Compose8(fAB, fBC, fCD, fDE, fEF, fFG, fGH, fHI), fIJ)
}

// Compose10 is a function that takes ten monadic functions and returns a single
// monadic function that represents the composition of the ten functions.
//
// This is a workaround for Go's type system which does not allow for dynamic
// function types in a list.
//
// If you need more than ten you can use the compose the Compose functions
// together. For example, `Compose(Compose5(f1, f2, f3, f4, f5), Compose5(f6,
// f7, f8, f9, f10))`
func Compose10[A, B, C, D, E, F, G, H, I, J, K any](
	fAB MonadicFunc[A, B],
	fBC MonadicFunc[B, C],
	fCD MonadicFunc[C, D],
	fDE MonadicFunc[D, E],
	fEF MonadicFunc[E, F],
	fFG MonadicFunc[F, G],
	fGH MonadicFunc[G, H],
	fHI MonadicFunc[H, I],
	fIJ MonadicFunc[I, J],
	fJK MonadicFunc[J, K],
) MonadicFunc[A, K] {
	return Compose(Compose9(fAB, fBC, fCD, fDE, fEF, fFG, fGH, fHI, fIJ), fJK)
}
