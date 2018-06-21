// authors: wangoo
// created: 2018-06-21
// field test

package inherit

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	a := A{
		F1: "a",
	}
	b := B{
		A:  a,
		F2: "b",
	}
	c := C{
		B:  b,
		F3: "c",
	}
	assert.Equal(t, "a", c.F1)
	assert.Equal(t, "b", c.F2)

	c.F1 = "d"
	c.F2 = "e"
	assert.Equal(t, "d", c.F1)
	assert.Equal(t, "e", c.F2)
}
