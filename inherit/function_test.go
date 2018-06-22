// authors: wangoo
// created: 2018-06-21
// function test

package inherit

import (
	"testing"
	"fmt"
)

// --------------------
type FA struct {
	f1 string
}

func (f FA) do1() {
	fmt.Println("f1", f.f1)
}

// --------------------
type FB struct {
	FA
	f2 string
}

func (f FB) do2() {
	fmt.Println("f2", f.f2)
}

// --------------------
func TestFu(t *testing.T) {
	a := FA{f1: "x"}
	b := FB{f2: "y"}

	a.do1()

	b.do1() // 调用父结构体的方法
	b.do2()

	b = FB{FA: a, f2: "z"}
	b.do1()
	b.do2()
}
