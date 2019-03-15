package basic

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

// Equals of struct, array, map
func TestEquals(t *testing.T) {
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
	u21 := &User{"u2", 18}
	u22 := u2
	assert.True(t, u2 != u21)
	assert.True(t, *u2 == *u21)
	assert.True(t, u2 == u22)

	arr1 := &[]User{u1, u11}
	arr11 := &[]User{u1, u11}
	arr12 := arr1
	assert.True(t, arr1 != arr11)
	assert.True(t, arr1 == arr12)

	m1 := &map[string]User{"1": u1, "2": u11}
	m11 := &map[string]User{"1": u1, "2": u11}
	m12 := m1
	assert.True(t, m1 != m11)
	assert.True(t, m1 == m12)

	v1 := reflect.ValueOf(u1)
	v11 := reflect.ValueOf(u1)
	assert.True(t, v1.Interface() == v11.Interface())
	assert.False(t, v1.CanAddr())

	v2 := reflect.ValueOf(u2)
	v21 := reflect.ValueOf(u2)
	p2 := v2.Pointer()
	p21 := v21.Pointer()
	assert.True(t, p2 == p21)

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
