// authors: wangoo
// created: 2018-06-22
// test function

package function

import (
	"fmt"
	"testing"
)

// ----------------------------------------
// 函数定义
func F1() {
	fmt.Println("f1")
}

// ----------------------------------------
type S1 struct {
}

// 函数定义struct对象上的方法
func (s S1) F1() {
	fmt.Println("S1.f1")
}

// ----------------------------------------
type S2 S1

// 函数定义自定义类型上的方法
func (s S2) F1() {
	fmt.Println("S2.f1")
}

// ----------------------------------------
type T1 string

// 函数定义自定义类型(string)上的方法, 不合法但编译没报错
func (s T1) F1() {
	fmt.Println("T1.f1")
}

// ----------------------------------------
type T2 int

// 函数定义自定义类型(int)上的方法, 不合法但编译没报错
func (s T2) F1() {
	fmt.Println("T2.f1")
}

// ----------------------------------------
type T3 bool

// 函数定义自定义类型(int)上的方法, 不合法但编译没报错
func (s T3) F1() {
	fmt.Println("T3.f1")
}

// ----------------------------------------

func TestFunc(t *testing.T) {
	F1()
	S1{}.F1()
	S2{}.F1()

	callT1("test")
	callT2(100)
	callT3(true)

}

func callT1(o interface{}) {
	if t1, ok := o.(T1); ok {
		t1.F1()
	} else {
		fmt.Println("can't call string T1.f1()")
	}
}

func callT2(o interface{}) {
	if t2, ok := o.(T2); ok {
		t2.F1()
	} else {
		fmt.Println("can't call int T2.f1()")
	}
}

func callT3(o interface{}) {
	if t3, ok := o.(T3); ok {
		t3.F1()
	} else {
		fmt.Println("can't call bool T3.f1()")
	}
}
