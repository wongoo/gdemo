package basic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSliceAppendNil(t *testing.T) {
	type Msg struct {
		Text string
	}
	a := Msg{Text: "a"}
	c := Msg{Text: "c"}

	var list []*Msg

	list = append(list, &a)
	list = append(list, nil) // append nil
	list = append(list, &c)

	assert.Equal(t, 3, len(list))
	assert.Equal(t, "a", list[0].Text)
	assert.Nil(t, list[1])
	assert.Equal(t, "c", list[2].Text)

	var list2 []*Msg
	v := reflect.ValueOf(list)
	vv := reflect.ValueOf(list2)

	for i := 0; i < v.Len(); i++ {
		vv = reflect.Append(vv, v.Index(i)) // support append nil value
	}
	list2 = vv.Interface().([]*Msg)

	assert.Equal(t, 3, len(list2))
	assert.Equal(t, "a", list2[0].Text)
	assert.Nil(t, list2[1])
	assert.Equal(t, "c", list2[2].Text)
}

func TestSliceOutOfIndex(t *testing.T) {
	s := make([]int, 3, 3)
	fmt.Println(s[3:]) // nil slice
	fmt.Println(s[4:]) // slice bounds out of range
}

type sscc struct {
	s string
}

func TestSliceCopyNotOverwrite(t *testing.T) {
	e := make([]sscc, 3)
	e[0].s = "a"
	e[1].s = "b"
	e[2].s = "c"

	c1 := e[0]
	assert.Equal(t, "a", c1.s)
	copy(e[0:], e[1:])
	assert.Equal(t, "a", c1.s)
	assert.Equal(t, "b", e[0].s)

	a := make([]*sscc, 3)
	a[0] = &sscc{s: "1"}
	a[1] = &sscc{s: "2"}
	a[2] = &sscc{s: "3"}

	c2 := a[0]
	assert.Equal(t, "1", c2.s)
	copy(a[0:], a[1:])
	assert.Equal(t, "1", c2.s)
	assert.Equal(t, "2", a[0].s)

}
