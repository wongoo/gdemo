package reflection

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// TestZeroDiff test whether the zero value is the same, including type and value
// The zero value of different types are different.
func TestZeroDiff(t *testing.T) {
	nilValue := reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem())
	assert.Equal(t, "<interface {} Value>", nilValue.String())
	assert.Equal(t, "interface {}", nilValue.Type().String())
	assert.Equal(t, nil, nilValue.Interface())
	assert.True(t, nilValue.IsNil())
	assert.True(t, nilValue.IsValid())

	type s struct {
		a  interface{}
		A1 interface{}
		b  *s
		B1 *s
	}

	x := s{}
	v := reflect.ValueOf(x)

	fieldA := v.FieldByName("a")
	assert.Equal(t, "<interface {} Value>", fieldA.String())
	assert.Equal(t, "interface {}", fieldA.Type().String())
	assert.NotEqual(t, nilValue, fieldA)
	assert.NotEqual(t, reflect.Zero(fieldA.Type()), fieldA)
	// assert.Equal(t, nil, fieldA.Interface()) // panic: cannot return value obtained from unexported field or method
	assert.True(t, fieldA.IsNil())
	assert.True(t, fieldA.IsValid())

	fieldA1 := v.FieldByName("A1")
	assert.Equal(t, nil, fieldA1.Interface())

	fieldB := v.FieldByName("b")
	assert.Equal(t, "<*reflection.s Value>", fieldB.String())
	assert.Equal(t, "*reflection.s", fieldB.Type().String())
	assert.NotEqual(t, nilValue, fieldB)
	assert.NotEqual(t, nilValue, fieldB.Elem())
	assert.NotEqual(t, reflect.Zero(fieldB.Type()), fieldB)
	// assert.Equal(t, nil, fieldB.Interface()) // panic: cannot return value obtained from unexported field or method
	assert.True(t, fieldB.IsNil())
	assert.True(t, fieldB.IsValid())

	fieldB1 := v.FieldByName("B1")
	assert.NotEqual(t, nil, fieldB1.Interface()) // <nil>(<nil>)  !=  *reflection.s((*reflection.s)(nil))
	assert.Equal(t, reflect.Zero(fieldB.Type()).Interface(), fieldB1.Interface())
}
