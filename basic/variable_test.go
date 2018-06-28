// authors: wangoo
// created: 2018-05-24
// variable demo

package basic

import (
	"fmt"
	"testing"
)

// define one const basic
const c1 = "v1"

// define multiple const variables
const (
	c2 = "v1"
	c3 = "v2"

	// define int variables
	c4 = 1
	c5 = 2
)

func TestVar(t *testing.T) {

	// print variables value
	fmt.Printf("c1=%v\n", c1)
	fmt.Printf("c2=%v\n", c2)
	fmt.Printf("c3=%v\n", c3)
	fmt.Printf("c4=%v\n", c4)
	fmt.Printf("c5=%v\n", c5)

	// define basic using var and assign
	var v1 = "x1"
	fmt.Println("v1:", v1)

	// define basic first
	var v2 string
	// assign
	v2 = "x2"
	fmt.Println("v2:", v2)

	// define and assign
	v3 := "x3"
	fmt.Println("v3:", v3)

	// multiple define and assign
	a, b, c := 1, 2, 3
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	// define multiple in parentheses
	var (
		x = 1
		y = 2
	)
	fmt.Printf("x=%v\n", x)
	fmt.Printf("y=%v\n", y)
}
