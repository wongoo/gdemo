// authors: wangoo
// created: 2018-06-22
// interfacc test

package interfacc

import (
	"testing"
	"fmt"
)

type Namer interface {
	Name() string
}

//-----------------------------
type S struct {
	f1 string
}

func (s S) Name() string {
	return "S"
}

//-----------------------------
type T struct {
}

func (t *T) Name() string {
	return "T"
}

//-----------------------------

// 演示func调用都是值传递,参数如果是struct对象则会复制一个对象，参数如果是指针这是传递指针的地址
// 如果方法是定义在struct对象上, 则对象和指针均可以访问此方法，且都可以被类型转换识别
func TestA(t *testing.T) {
	s1 := S{}
	fmt.Printf("s1 name=%v, addr=%p pointer addr=%p\n", s1.Name(), s1, &s1)
	p(s1)
	p(&s1)

	fmt.Println("----------")

	s2 := &S{}
	fmt.Printf("s2 name=%v, addr=%p pointer addr=%p\n", s2.Name(), s2, &s2)
	p(*s2)
	p(s2)
}

// 演示struct作为参数传递，func内部对值的修改不会影响到func外部原始struct对象
func TestB(t *testing.T) {
	s1 := S{f1: "x"}
	fmt.Printf("s1 f1=%v\n", s1.f1)
	update(s1)
	fmt.Printf("s1 after update(b), f1=%v\n", s1.f1)

	fmt.Println("----------")

	s2 := S{f1: "x"}
	fmt.Printf("s2 f1=%v\n", s2.f1)
	update(&s2)
	fmt.Printf("s2 after update(&b), f1=%v\n", s2.f1)
}

// 演示方法定义在 <struct指针对象> 上的差异，对象和指针都可以访问此方法，但只有指针可以被类型转换识别
func TestT(t *testing.T) {
	t1 := T{}
	fmt.Printf("t1 name=%v, addr=%p, pointer addr=%p\n", t1.Name(), t1, &t1)
	p(t1)

	fmt.Println("----------")

	t2 := &T{}
	fmt.Printf("t2 name=%v, addr=%p, pointer addr=%p\n", t2.Name(), t2, &t2)
	p(t2)
}

func p(o interface{}) {
	if n, ok := o.(Namer); ok {
		fmt.Printf("--> its namer, name=%v, addr=%p, pointer addr=%p\n", n.Name(), o, &o)
	} else {
		fmt.Printf("--> not namer, addr=%p, pointer addr=%p\n", o, &o)
	}
}

func update(o interface{}) {
	if s, ok := o.(S); ok {
		s.f1 = "y"
		fmt.Printf("--> struct S, change f1=%v\n", s.f1)
	}
	if s, ok := o.(*S); ok {
		s.f1 = "y"
		fmt.Printf("--> struct *S, change f1=%v\n", s.f1)
	}
}
