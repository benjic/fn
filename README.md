# Fn

This is a collection of functional types that allow you to build type safety
into how code is composed.

## Optional

The optional package provide a way to encapsulate a series of computations that
depend on a value. If at any point a value is no possible an empty optional
can be returned.

This is useful in Go in that in helps you protect values by avoiding the `nil`
case.

```go
// bad
func getNonZeroValue(val int) *int {
    if val < 0 { return nil }

    return val
}

// good
func getNonZeroOptional(val int) optional.Optional[int] {
    if val < 0 { return Empty() }

    return optional.Value(val)
}

func main() {

    // forgot to nil check so dereference panics!
    sum(&getNonZeroValue(-1), 1)


    // With the optional I _must_ use a function with an error at the calling
    // site to get a value. I do not need to return errors from my pure function
    // that gets a nonzero value.
    val, err := optional.Must(getNonZeroOptional(-1))
    if err != {
        fmt.Printf("invalid value: %w", err)
    }

    sum(val, 1)
}
```

This pattern allows us to use a _monad_ pattern to compose computations and
pass an optional value through. I do not need a to have functions that rely
on multiple return values _easing composition._

```go
func getNonZeroOptional(val int) optional.Optional[int] {
    if val < 0 { return Empty() }

    return optional.Value(val)
}

func main() {
    // Define a series of steps that return empty values and turns them
    // to a string.
    getNonZeroString := optional.Compose(
        getNonZeroOptional,

        // Sometimes our functions do not use the optional type so `Lift`
        // allows us to "lift" the pure function into the optional
        // monad so it can be executed in the context of an `Optional`.
        optional.Lift(
            func(v int) string { return fmt.Sprintf("%d", v)}
        )
    )

    val, err := getNonZeroString(rand.Intn(21) - 10))
    if err != nil {
        fmt.Printf("non zero value found")
        return
    }

    fmt.Printf("got %s", val)
}
```

This allows us to avoid using errors to communicate to our caller an unusable
value was part of the call stack.
