// authors: wangoo
// created: 2018-06-29
// test s2

package main

import (
	"fmt"
	"github.com/wongoo/gdemo/scope/s1"
	"github.com/wongoo/gdemo/scope/core"
)

func main() {
	m := s1.GetM()
	fmt.Println(m.Name)

	m.Name = "update"

	fmt.Println(m.Name)

	fmt.Println(GetName(m))
	fmt.Println(GetDesc(m))
}

func GetName(m *core.M) string {
	return m.Name
}
func GetDesc(m *core.M) string {
	return m.Desc()
}
