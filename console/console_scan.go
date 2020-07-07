package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	quickSort(a, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
}

func quickSort(a []int, l, r int) {
	if l >= r {
		return
	}

	x := a[l]
	i := l + 1
	j := r

	for {

		for i < r && a[i] < x {
			i++
		}

		for j > l && a[j] > x {
			j--
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	a[l], a[j] = a[j], a[l]

	quickSort(a, l, j-1)
	quickSort(a, j+1, r)
}
