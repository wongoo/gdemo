// authors: wangoo
// created: 2018-06-22
//

package main

import (
	"fmt"
	"github.com/wongoo/gdemo/interfacc/p1"
)

type T struct {
}

func (t T) Hello() {
	fmt.Println("hello")
}

func (t T) hi() {
	fmt.Println("hi")
}

func callHello(t interface{}) {
	if h, ok := t.(p1.TA); ok {
		h.Hello()
	} else {
		fmt.Println("not p1.TA")
	}
}

func main() {
	t := T{}
	fmt.Println(t)
	callHello(t)
}
