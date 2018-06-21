// authors: wangoo
// created: 2018-06-21
// inherit

package inherit

type A struct {
	F1 string
}

type B struct {
	A
	F2 string
}

type C struct {
	B
	F3 string
}
