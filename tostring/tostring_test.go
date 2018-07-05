// authors: wangoo
// created: 2018-07-05
// tostring test

package tostring

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

type TS1 struct {
}

func (t *TS1) String() string {
	return "I'am TS1"
}

type TS2 struct {
}

func (t *TS2) Error() string {
	return "TS2 Error"
}

func (t *TS2) String() string {
	return "I'am TS2"
}

func TestToString(t *testing.T) {
	ts1 := &TS1{}
	assert.Equal(t, "I'am TS1", fmt.Sprint(ts1))

	ts2 := &TS2{}
	assert.Equal(t, "TS2 Error", fmt.Sprint(ts2))
}
