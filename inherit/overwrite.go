// authors: wangoo
// created: 2018-06-21
// interface

package inherit

// ----------------------
type S1 struct {
	F1 string
}

func (s S1) Name() string {
	return s.F1
}

// ----------------------
type S2 struct {
	S1
	F2 string
}

// ----------------------
type S3 struct {
	S1
	F3 string
}

func (s S3) Name() string {
	return s.F3
}
