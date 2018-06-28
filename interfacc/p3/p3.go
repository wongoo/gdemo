// authors: wangoo
// created: 2018-06-22
// define some interface

package p3

type TA interface {
	Hello()
}

type TB interface {
	TA
	hi()
}
