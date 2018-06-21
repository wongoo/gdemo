// authors: wangoo
// created: 2018-06-21
// function test

package inherit

import "testing"

func TestFu(t *testing.T) {
	a := FA{f1: "x"}
	b := FB{f2: "y"}

	a.do1()

	b.do1()
	b.do2()

	b = FB{FA: a, f2: "z"}
	b.do1()
	b.do2()
}
