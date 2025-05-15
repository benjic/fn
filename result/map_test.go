package result

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	t.Run("applies func to value", func(t *testing.T) {
		o := Map(
			Value(42),
			func(i int) string { return fmt.Sprintf("%d", i) },
		)

		val, err := Must(o)
		assert.NoError(t, err)
		assert.Equal(t, "42", val)
	})

	t.Run("skips err", func(t *testing.T) {
		testErr := fmt.Errorf("test error")
		o := Map(
			Error[int](testErr),
			func(i int) string { return fmt.Sprintf("%d", i) },
		)

		val, err := Must(o)
		assert.ErrorIs(t, err, testErr)
		assert.Equal(t, "", val)
	})

}
