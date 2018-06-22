// authors: wangoo
// created: 2018-06-21
// test interface

package inherit

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// ----------------------
type S1 struct {
	F1 string
}

func (s S1) Name() string {
	return s.F1
}

// ----------------------
type S2 struct {
	S1
	F2 string
}

// ----------------------
type S3 struct {
	S1
	F3 string
}

// S3的Name方法将覆盖父结构体S1的Name方法
func (s S3) Name() string {
	return s.F3
}

// ----------------------
func TestOverwrite(t *testing.T) {
	s1 := S1{
		F1: "x",
	}
	assert.Equal(t, "x", s1.Name())

	s2 := S2{
		S1: s1,
		F2: "y",
	}
	// 调用s2.Name()和s2.S1.Name()是等效的
	assert.Equal(t, "x", s2.Name())
	assert.Equal(t, "x", s2.S1.Name())

	s3 := S3{
		S1: s1,
		F3: "z",
	}
	assert.Equal(t, "z", s3.Name())    // 调用s3自己覆盖的方法Name()
	assert.Equal(t, "x", s3.S1.Name()) //调用父结构体S1的方法Name()
}
