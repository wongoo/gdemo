// authors: wangoo
// created: 2018-06-29
// core

package core

type M struct {
	Name string
}

func (m *M) Desc() string {
	return m.Name
}
