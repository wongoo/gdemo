package reflection

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestStructPointerChange(t *testing.T) {
	type x struct {
		name string
	}

	x1 := x{"abc"}
	t.Logf("x1 address: %p, %p", &x1, &x1.name)

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&x1.name))
	t.Logf("name string addr: %d, %d", hdr.Len, hdr.Data)

	p1 := &x1
	t.Logf("p1 address: %p", p1) // 地址一致

	v1 := reflect.ValueOf(x1)
	// panic: reflect: call of reflect.Value.Pointer on struct Value
	// t.Logf("v1 address: %p", unsafe.Pointer(v1.Pointer()))

	i1 := v1.Interface()
	t.Logf("i1 address: \t %p <---- 地址变了,interface()会创建新的结构体", &i1) // 地址改变了, 只有指针操作才能保持引用

	x1copy := i1.(x)
	t.Logf("copied address:\t %p <---- 地址变了，类型转换也会创建新的结构体", &x1copy) // 地址改变了, 只有指针操作才能保持引用
	hdr = (*reflect.StringHeader)(unsafe.Pointer(&x1copy.name))
	t.Logf("copy string addr: %d, %d <---- 虽然是新的结构体，但内部name的运行时值还是只有一份", hdr.Len, hdr.Data)

	// 指针反射创建新value
	pv1 := reflect.New(v1.Type())
	pv1.Elem().Set(v1)
	t.Logf("pv1 address:\t %p <---- 地址变了，反射set的会再创建新的结构体", unsafe.Pointer(pv1.Pointer())) //  地址改变了, 只有指针操作才能保持引用

	vp1 := reflect.ValueOf(p1)
	t.Logf("vp1 address: %p", unsafe.Pointer(vp1.Pointer())) // 地址一致
	ip1 := vp1.Interface()
	t.Logf("ip1 address: %p", ip1) // 地址一致

	// 指针反射创建新value
	ppv1 := reflect.New(vp1.Type())
	ppv1.Elem().Set(vp1)
	t.Logf("ppv1 address: %p", unsafe.Pointer(ppv1.Elem().Pointer())) // 地址一致

}
