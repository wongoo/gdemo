// authors: wangoo
// created: 2018-06-21
// test interface

package inherit

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOverwrite(t *testing.T) {
	s1 := S1{
		F1: "x",
	}
	assert.Equal(t, "x", s1.Name())

	s2 := S2{
		S1: s1,
		F2: "y",
	}
	assert.Equal(t, "x", s2.Name())
	assert.Equal(t, "x", s2.S1.Name())

	s3 := S3{
		S1: s1,
		F3: "z",
	}
	assert.Equal(t, "z", s3.Name())
	assert.Equal(t, "x", s3.S1.Name())
}
