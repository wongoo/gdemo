package unsafego

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

func TestUnsafeSize(t *testing.T) {

	var s string
	t.Log("Alignof string : ", unsafe.Alignof(s))
	t.Log("Sizeof  string : ", unsafe.Sizeof(s))

	var b bool
	t.Log("Alignof bool : ", unsafe.Alignof(b))
	t.Log("Sizeof  bool : ", unsafe.Sizeof(b))

	var bt byte
	t.Log("Alignof byte : ", unsafe.Alignof(bt))
	t.Log("Sizeof  byte : ", unsafe.Sizeof(bt))

	var i8 int8
	t.Log("Alignof int8 : ", unsafe.Alignof(i8))
	t.Log("Sizeof  int8 : ", unsafe.Sizeof(i8))

	var i16 int16
	t.Log("Alignof int16 : ", unsafe.Alignof(i16))
	t.Log("Sizeof  int16 : ", unsafe.Sizeof(i16))

	var i32 int32
	t.Log("Alignof int32 : ", unsafe.Alignof(i32))
	t.Log("Sizeof  int32 : ", unsafe.Sizeof(i32))

	var i64 int64
	t.Log("Alignof int64 : ", unsafe.Alignof(i64))
	t.Log("Sizeof  int64 : ", unsafe.Sizeof(i64))

	var i int
	t.Log("Alignof int : ", unsafe.Alignof(i))
	t.Log("Sizeof  int : ", unsafe.Sizeof(i))

	var ui32 uint32
	t.Log("Alignof uint32 : ", unsafe.Alignof(ui32))
	t.Log("Sizeof  uint32 : ", unsafe.Sizeof(ui32))

	var ui64 uint64
	t.Log("Alignof uint64 : ", unsafe.Alignof(ui64))
	t.Log("Sizeof  uint64 : ", unsafe.Sizeof(ui64))

	var f32 float32
	t.Log("Alignof float32 : ", unsafe.Alignof(f32))
	t.Log("Sizeof  float32 : ", unsafe.Sizeof(f32))

	var f64 float64
	t.Log("Alignof float64 : ", unsafe.Alignof(f64))
	t.Log("Sizeof  float64 : ", unsafe.Sizeof(f64))

	type t1 struct {
		b   byte
		i32 int32
		i16 int16
		i64 int64
		f32 float32
		f64 float64
	}

	tt := t1{}
	t.Log("Alignof t1 : ", unsafe.Alignof(tt))
	t.Log("Sizeof  t1 : ", unsafe.Sizeof(tt))

	t.Log("Offsetof of b : ", unsafe.Offsetof(tt.b))
	t.Log("Offsetof of i32 : ", unsafe.Offsetof(tt.i32))
	t.Log("Offsetof of i16 : ", unsafe.Offsetof(tt.i16))
	t.Log("Offsetof of i64 : ", unsafe.Offsetof(tt.i64))
	t.Log("Offsetof of f32 : ", unsafe.Offsetof(tt.f32))
	t.Log("Offsetof of f64 : ", unsafe.Offsetof(tt.f64))
}

func TestUnsafeStructSize(t *testing.T) {
	type t1 struct {
		b   bool
		i16 int16
		i32 int32
	}
	type t2 struct {
		b   bool
		i32 int32
		i16 int16
	}

	x1 := t1{}
	x2 := t2{}
	t.Logf("t1 Alignof: %d, Sizeof: %d", unsafe.Alignof(x1), unsafe.Sizeof(x1))
	t.Logf("t2 Alignof: %d, Sizeof: %d", unsafe.Alignof(x2), unsafe.Sizeof(x2))

	t.Logf("t1 Offsetof b: %d, i16: %d , i32: %d", unsafe.Offsetof(x1.b), unsafe.Offsetof(x1.i16), unsafe.Offsetof(x1.i32))
	t.Logf("t2 Offsetof b: %d , i32: %d, i16: %d", unsafe.Offsetof(x2.b), unsafe.Offsetof(x2.i32), unsafe.Offsetof(x2.i16))
}

func TestPointerConversion(t *testing.T) {
	var f64 float64 = 1.23
	t.Logf("f64: %f", f64)
	ip64 := (*uint64)(unsafe.Pointer(&f64))
	i64 := *ip64
	t.Logf("ip64: %p, value: %d", ip64, *ip64)
	t.Logf("i64: %d, addr: %p", i64, &i64)

	var i32 int32 = 234
	t.Logf("i32: %d, addr: %p", i32, &i32)
	ip64 = (*uint64)(unsafe.Pointer(&i32))
	i64 = *ip64
	t.Logf("ip64: %p, value: %d", ip64, *ip64)
	t.Logf("i64: %d, addr: %p", i64, &i64)

	i32 = 345
	t.Logf("i32: %d, addr: %p", i32, &i32)
	t.Logf("ip64: %p, value: %d", ip64, *ip64)
	t.Logf("i64: %d, addr: %p", i64, &i64)
}

func TestStructPointerConversion(t *testing.T) {
	type x struct {
		f1 int
		f2 int
	}

	i1 := 123
	i2 := 456

	t.Logf("i1 addr: %p, value: %d", &i1, i1)
	t.Logf("i2 addr: %p, value: %d", &i2, i2)

	xp := (*x)(unsafe.Pointer(&i1))
	t.Logf("xp addr: %p, value: %v", xp, xp)

	x1 := *xp
	t.Logf("x1 addr: %p, value: %v", &x1, x1)
}

func TestStructFieldPointerConversion(t *testing.T) {
	type x struct {
		f1 int
		f2 *int
	}
	i1 := 123
	i2 := 456

	t.Logf("i1 addr: %p, value: %d", &i1, i1)
	t.Logf("i2 addr: %p, value: %d", &i2, i2)

	xp := (*x)(unsafe.Pointer(&i1))
	t.Logf("xp addr: %p, value: %v", xp, xp)

	x1 := *xp
	t.Logf("x1 addr: %p, value: %v", &x1, x1)

	t.Log(*x1.f2)
}

func TestUnsafeReflect(t *testing.T) {
	a := "hello"
	b := "world"
	t.Logf("a addr: %p,   value: %s", &a, a)
	t.Logf("b addr: %p,   value: %s", &b, b)

	var p *string
	t.Logf("p addr: %p", p)

	p = &a
	t.Logf("p addr: %p", p)

	sh1 := (*reflect.StringHeader)(unsafe.Pointer(p))
	sh2 := (*reflect.StringHeader)(unsafe.Pointer(&b))
	s1 := (*string)(unsafe.Pointer(sh1))
	s2 := (*string)(unsafe.Pointer(sh2))

	t.Logf("sh1 addr: %p, value: %v", sh1, sh1)
	t.Logf("sh2 addr: %p, value: %v", sh2, sh2)
	t.Logf("s1 addr: %p, value: %s", s1, *s1)
	t.Logf("s2 addr: %p, value: %s", s2, *s2)

	sh1.Data = uintptr(unsafe.Pointer(&b))
	t.Logf("sh1 addr: %p, value: %v", sh1, sh1)
	t.Logf("s1 addr: %p,  value: %s", s1, *s1)

	sh1.Data = sh2.Data
	sh1.Len = sh2.Len
	t.Logf("s1 addr: %p,  value: %s", s1, *s1)
	t.Logf("a addr: %p, value: %s", &a, a)
	t.Logf("b addr: %p, value: %s", &b, b)

}

func TestPointerEquals(t *testing.T) {
	type x struct {
		f1 int
		f2 int
	}

	x1 := x{1, 2}

	p1 := unsafe.Pointer(&x1)
	p2 := unsafe.Pointer(&x1)

	assert.True(t, p1 == p2)

	p3 := unsafe.Pointer(&x1.f1)
	assert.True(t, p1 == p3)
}
