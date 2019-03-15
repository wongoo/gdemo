package reflection_test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func logTypeName(t *testing.T, typ reflect.Type) {
	t.Log(typ, " : ", typ.Name())
}
func TestTypeName(t *testing.T) {
	logTypeName(t, reflect.TypeOf([]int{}))
	logTypeName(t, reflect.TypeOf([][]int{}))
	logTypeName(t, reflect.TypeOf([]string{}))
	logTypeName(t, reflect.TypeOf(map[string]string{}))

	type TT struct {
		F1 []string
	}
	typ := reflect.TypeOf(TT{})
	logTypeName(t, typ)
	logTypeName(t, reflect.TypeOf([]TT{}))
	logTypeName(t, reflect.TypeOf(map[string]TT{}))

	if f1, ok := typ.FieldByName("F1"); ok {
		logTypeName(t, f1.Type)
	}
}

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

func TestPointerValueSet(t *testing.T) {
	type User struct {
		name string
		age  int
	}

	u1 := User{"u1", 10}
	p1 := &u1
	p2 := &p1
	p3 := &p2
	fmt.Println(p1, p2, p3)

	v1 := reflect.New(reflect.TypeOf(p1))
	// v1.Set(reflect.ValueOf(u1)) //error
	// v1.Set(reflect.ValueOf(p1)) //error
	v1.Elem().Set(reflect.ValueOf(p1))

	v2 := reflect.New(reflect.TypeOf(p2))
	// v2.Elem().Set(reflect.ValueOf(p1)) //error
	v2.Elem().Set(reflect.ValueOf(p2))

	v3 := reflect.New(reflect.TypeOf(p3))
	v3.Elem().Set(reflect.ValueOf(p3))
}

func TestAddr(t *testing.T) {
	type User struct {
		name string
		age  int
	}

	u1 := User{"u1", 18}
	u11 := User{"u1", 18}
	u12 := User{"u1", 20}

	assert.True(t, u1 == u11)
	assert.True(t, u1 != u12)

	user2 := User{"u2", 18}
	u2 := &user2

	v1 := reflect.ValueOf(u1)
	v11 := reflect.ValueOf(u1)
	assert.True(t, v1.Interface() == v11.Interface())
	assert.False(t, v1.CanAddr())

	v2 := reflect.ValueOf(u2)
	v21 := reflect.ValueOf(u2)
	// Pointer 相当于获得指针值作为地址
	p2 := v2.Pointer()
	p21 := v21.Pointer()
	assert.True(t, p2 == p21)

	// UnsafeAddr 相当于 取地址操作 `&`
	addr2 := v2.Elem().UnsafeAddr()
	addr21 := v21.Elem().UnsafeAddr()
	assert.True(t, addr2 == addr21)
	assert.True(t, p2 == addr2)

	userType := reflect.TypeOf(user2)
	userPtrType := reflect.PtrTo(userType)
	t.Logf("user type: %v , ptr type: %v", userType, userPtrType)

	v22 := reflect.NewAt(userPtrType, unsafe.Pointer(u2))
	p22 := v22.Pointer()
	assert.True(t, p2 == p22)

	user3 := reflect.New(userType)
	p3 := reflect.New(userPtrType)
	p3.Elem().Set(user3)
	t.Logf("user3 type: %v , p3 type: %v", user3.Type(), p3.Type())

	assert.True(t, user3.Interface() == p3.Elem().Interface())
	assert.True(t, p3.Pointer() == p3.Elem().UnsafeAddr())
	assert.True(t, p3.Elem().Pointer() == user3.Elem().UnsafeAddr())


}
