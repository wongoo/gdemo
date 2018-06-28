// authors: wangoo
// created: 2018-06-28
// method test

package method

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

// ------------------------
type X struct {
	Name string
}

func (x X) GetName() string {
	fmt.Printf("get name of x %p\n", &x)
	return x.Name
}
func (x X) SetName(name string) {
	fmt.Printf("set name of x %p\n", &x)
	x.Name = name
}

// ------------------------
type Y struct {
	Name string
}

func (y *Y) GetName() string {
	fmt.Printf("get name of y %p\n", y)
	return y.Name
}
func (y *Y) SetName(name string) {
	fmt.Printf("get name of y %p\n", y)
	y.Name = name
}

// ------------------------

func TestX(t *testing.T) {
	x1 := X{}
	fmt.Printf("new x %p\n", &x1)
	x1.SetName("x1")

	assert.Equal(t, "", x1.GetName())

	x2 := &X{}
	fmt.Printf("new x %p\n", x2)
	x2.SetName("x2")

	assert.Equal(t, "", x2.GetName())
}

func TestY(t *testing.T) {
	y1 := Y{}
	fmt.Printf("new y %p\n", &y1)
	y1.SetName("y1")

	assert.Equal(t, "y1", y1.GetName())

	y2 := &Y{}
	fmt.Printf("new y %p\n", y2)
	y2.SetName("y2")

	assert.Equal(t, "y2", y2.GetName())
}
