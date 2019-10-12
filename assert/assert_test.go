package assert

import "testing"

func TestAssert(t *testing.T) {
	assert := New(t)
	assert.Equal(1, 1)
	assert.NEqual(2, 1)
	assert.True(true)
	assert.False(false)
}
