// authors: wangoo
// created: 2018-06-21
// function

package inherit

import "fmt"

type FA struct {
	f1 string
}

func (f FA) do1() {
	fmt.Println("f1", f.f1)
}

type FB struct {
	FA
	f2 string
}

func (f FB) do2() {
	fmt.Println("f2", f.f2)
}
