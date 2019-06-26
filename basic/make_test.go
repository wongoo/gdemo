package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeArray(t *testing.T) {
	c := make(chan struct{}, 5)
	assert.Equal(t, 0, len(c))

	c <- struct{}{}
	assert.Equal(t, 1, len(c))

	c1 := make([]chan struct{}, 5)
	assert.Equal(t, 5, len(c1))
}
