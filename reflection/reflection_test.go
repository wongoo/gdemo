package reflection_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	type A struct{}

	a := A{}
	ap := &a
	typA := reflect.TypeOf(a)   // A
	typAp := reflect.TypeOf(ap) // *A

	t.Log("struct type:", typA)                      // A
	t.Log("struct pointer type:", typAp)             // *A
	t.Log("struct pointer type kind:", typAp.Kind()) // ptr
	t.Log("struct pointer type elem:", typAp.Elem()) // A

	assert.Equal(t, typA, typAp.Elem())

	valueA := reflect.ValueOf(a)   // value of A
	valueAp := reflect.ValueOf(ap) // value of *A
	t.Log("value:", valueA)
	t.Log("value of pointer:", valueAp)
	t.Log("value kind of pointer:", valueAp.Kind()) //ptr

	valueA2 := valueAp.Elem() // value of A from value of *A
	t.Log("value from pointer:", valueA2)

	// assert.Equal(t, valueA, valueA2) //NOT equal

	x := valueA.Interface()
	a2, ok := x.(A)
	assert.True(t, ok)
	assert.Equal(t, a, a2)
	x2 := valueA2.Interface() // valueA and valueA2 ref to the same instance A
	assert.Equal(t, x, x2)

	y := valueAp.Interface()
	ap2, ok := y.(*A)
	assert.True(t, ok)
	assert.Equal(t, ap, ap2)

}
