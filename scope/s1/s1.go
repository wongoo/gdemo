// authors: wangoo
// created: 2018-06-29
// s1

package s1

import "github.com/wongoo/gdemo/scope/core"

var m *core.M

func init() {
	m = &core.M{
		Name: "test",
	}
}

func GetM() *core.M {
	return m
}
