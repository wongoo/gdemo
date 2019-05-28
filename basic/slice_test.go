package basic

import (
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
