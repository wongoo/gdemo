// authors: wangoo
// created: 2018-06-22
// TODO add description about this file

package p3

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

type TAS struct {
}

func (t TAS) Hello() {
	fmt.Println("hello")
}
func (t TAS) hi() {
	fmt.Println("hi")
}

func TestTA(t *testing.T) {
	tas := &TAS{}
	assert.True(t, isTA(tas))
	assert.True(t, isTB(tas))
}

func isTA(o interface{}) bool {
	_, ok := o.(TA)
	return ok
}
func isTB(o interface{}) bool {
	_, ok := o.(TB)
	return ok
}