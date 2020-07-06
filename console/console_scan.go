package main

import (
	"fmt"
)

func main() {
	n := 0

	_, _ = fmt.Scanln(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&a[i])
	}

	fmt.Println(a)
}
