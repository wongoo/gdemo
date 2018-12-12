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

func TestValueSet(t *testing.T) {
	type B struct {
		s string
		R int
	}

	b := B{s: "x", R: 1}
	bp := &b

	bType := reflect.TypeOf(b)
	bpType := reflect.TypeOf(bp)

	bValue := reflect.ValueOf(b)
	bpValue := reflect.ValueOf(bp)

	t.Log("b:", b)
	t.Log("b type:", bType)
	t.Log("b value:", bValue)
	t.Log("settability of b value:", bValue.CanSet())
	t.Log("bp:", bp)
	t.Log("bp type:", bpType)
	t.Log("bp value(1):", bpValue)
	t.Log("settability of bp value:", bpValue.CanSet())
	t.Log("settability of bp elem value:", bpValue.Elem().CanSet())

	b2 := B{s: "y", R: 2}

	bpValue.Elem().Set(reflect.ValueOf(b2))
	t.Log("bp value(2):", bpValue)

	newBPtrValue := reflect.New(bType)
	newBValue := newBPtrValue.Elem()
	//newBValue.FieldByName("s").SetString("z") // reflect.Value.SetString using value obtained using unexported field
	newBValue.FieldByName("R").SetInt(3)
	bpValue.Elem().Set(newBValue)
	t.Log("bp value(3):", bpValue)

	// bpValue.Elem().Set(reflect.New(bpType).Elem().Elem()) // panic: reflect: call of reflect.Value.Set on zero Value<Paste>
}
