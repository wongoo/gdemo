// authors: wangoo
// created: 2018-06-22
// test interface

package main

import (
	"fmt"
	"github.com/wongoo/gdemo/interfacc/p3"
)

type TA2 struct {
}

func (t TA2) Hello() {
	fmt.Println("hello")
}
func (t TA2) hi() () {
	fmt.Println("hi")
}

func main() {
	ta2 := &TA2{}
	callHello(ta2)
	// callHi(ta2) --> compile error

	fmt.Println(isTA(ta2)) // true
	fmt.Println(isTB(ta2)) // false
}

func isTA(o interface{}) bool {
	_, ok := o.(p3.TA)
	return ok
}

func isTB(o interface{}) bool {
	_, ok := o.(p3.TB)
	return ok
}

func callHello(h interface{ p3.TA }) {
	h.Hello()
}
func callHi(h interface{ p3.TB }) {
	h.Hello()
}
