package basic

import (
	"fmt"
	"testing"
)

var itemExists = struct{}{}

func TestNilStruct(t *testing.T) {
	doTestNilStruct()
	doTestNilStruct()
}

func doTestNilStruct() {
	fmt.Println("---------------")
	m := make(map[interface{}]struct{})
	a := struct{}{}
	b := struct{}{}

	fmt.Printf("item=%p\n", &itemExists)

	m[a] = itemExists
	m[b] = itemExists
	m[struct{}{}] = itemExists
	m["k1"] = itemExists
	m["k2"] = itemExists
	m["k3"] = struct{}{}

	fmt.Printf("a=%p\n", &a)
	fmt.Printf("b=%p\n", &b)

	for key, value := range m {
		fmt.Printf("key=%v, %p    value=%v, %p\n", key, &key, value, &value)
	}
}
