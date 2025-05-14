package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Value(t *testing.T) {
	// Test with a non-nil value
	opt := Value(42)
	assert.True(t, HasValue(opt))
	assert.False(t, IsEmpty(opt))

	val, err := Must(opt)
	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func Test_Empty(t *testing.T) {
	// Test with a non-nil value
	opt := Empty[int]()
	assert.False(t, HasValue(opt))
	assert.True(t, IsEmpty(opt))

	val, err := Must(opt)
	assert.ErrorIs(t, err, ErrMissingValue)
	assert.EqualError(t, err, "Optional[int]: optional value required")
	assert.Equal(t, 0, val, "empty value")
}
