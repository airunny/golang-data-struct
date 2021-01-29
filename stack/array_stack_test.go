package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArrayStack(t *testing.T) {
	s := NewArrayStack()
	assert.Nil(t, s.Top())

	//
	assert.Equal(t, s.Push(1), 0)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 1, s.Top())
	assert.Equal(t, 1, s.Pop())
	assert.True(t, s.IsEmpty())

	// again
	assert.Equal(t, s.Push(1), 0)
	assert.Equal(t, 1, s.Top())
	assert.False(t, s.IsEmpty())

	assert.Equal(t, s.Push(2), 0)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 2, s.Top())
	assert.Equal(t, 2, s.Pop())
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 1, s.Top())
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 1, s.Pop())
	assert.True(t, s.IsEmpty())
}
