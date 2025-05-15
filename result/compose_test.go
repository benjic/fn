package result

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bind(t *testing.T) {
	testErr := fmt.Errorf("test error")

	o := Bind(
		Bind(
			Bind(
				Value(42),
				Lift(func(i int) int {
					return i - 35
				}),
			),
			Lift(func(i int) string {
				return fmt.Sprintf("%d", i)
			}),
		),
		func(i string) Result[struct{}] {
			return Error[struct{}](testErr)
		},
	)

	val, err := Must(o)
	assert.ErrorIs(t, err, testErr)
	assert.Equal(t, struct{}{}, val)
}

func Test_Compose(t *testing.T) {
	errLessThan := fmt.Errorf("value less than")

	subtract := func(val int) func(int) int { return func(i int) int { return i - val } }

	skipLessThan := func(val int) func(int) Result[int] {
		return func(i int) Result[int] {
			if i < val {
				return Error[int](fmt.Errorf("%w: %d", errLessThan, val))
			}

			return Value(i)
		}
	}

	toString := func(i int) string { return fmt.Sprintf("%d", i) }

	append := func(s string) func(string) string { return func(i string) string { return i + s } }

	t.Run("successful bind", func(t *testing.T) {

		compute := Compose10(
			Lift(subtract(35)),
			skipLessThan(10),
			Lift(toString),
			Lift(append("t")),
			Lift(append("e")),
			Lift(append("s")),
			Lift(append("t")),
			Lift(append("")),
			Lift(append("")),
			Lift(append("")),
		)

		t.Run("early escape", func(t *testing.T) {
			// This results in a zero value
			val, err := Must(compute(42))
			assert.ErrorIs(t, err, errLessThan)
			assert.Equal(t, "", val)
		})

		t.Run("complete", func(t *testing.T) {
			val, err := Must(compute(50))
			assert.NoError(t, err, errLessThan)
			assert.Equal(t, "15test", val)
		})
	})

}
