package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	// int
	assert.Equal(t, true, ContainsInt([]int{1, 2, 3}, 1))
	assert.Equal(t, false, ContainsInt([]int{1, 2, 3}, 9))

	// int64
	assert.Equal(t, true, ContainsInt64([]int64{1, 2, 3}, 1))
	assert.Equal(t, false, ContainsInt64([]int64{1, 2, 3}, 9))

	// float64
	assert.Equal(t, true, ContainsFloat64([]float64{1.1, 2.3, 3}, 1.1))
	assert.Equal(t, false, ContainsFloat64([]float64{1, 2, 3.5}, 4.4))

	// string
	assert.Equal(t, true, ContainsString([]string{"a", "b", "c"}, "a"))
	assert.Equal(t, false, ContainsString([]string{"a", "b", "c"}, "d"))
}
