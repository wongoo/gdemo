package reflection

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

type ss struct {
	i1 int8
	i2 int8
	i3 int8
}

func TestConvertInt(t *testing.T) {
	s := &ss{
		i1: 1,
		i2: 2,
		i3: 4,
	}

	assert.Equal(t, int8(1), s.i1)
	assert.Equal(t, int8(2), s.i2)
	assert.Equal(t, int8(4), s.i3)

	t32 := *(*int32)(unsafe.Pointer(&s.i1))
	assert.Equal(t, int32(-1073479167), t32)
}
