// authors: wangoo
// created: 2018-06-21
// field test

package inherit

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// --------------------
type A struct {
	F1 string
}

// --------------------
type B struct {
	A // 继承了A的字段
	F2 string
}

// --------------------
type C struct {
	B // 继承了B的字段
	F3 string
}

// --------------------
func TestF(t *testing.T) {
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
	assert.Equal(t, "a", c.F1) //访问父结构体A的字段
	assert.Equal(t, "b", c.F2) //访问父结构体B的字段

	c.F1 = "d" //修改父结构体A的字段值
	c.F2 = "e" //修改父结构体B的字段值
	assert.Equal(t, "d", c.F1)
	assert.Equal(t, "e", c.F2)
}
