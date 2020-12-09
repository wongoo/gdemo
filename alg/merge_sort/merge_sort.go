package main

import (
	"fmt"
)

func main() {
	//var n int
	//fmt.Scanf("%d\n", &n)
	//a:=make([]int, n)
	//for i:=0;i<n;i++{
	//	fmt.Scanf("%d", &a[i])
	//}

	n := 5
	a := []int{3, 1, 2, 4, 5}

	temp := make([]int, n)

	mergeSort(a, temp, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
}

func mergeSort(a, temp []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) >> 1
	mergeSort(a, temp, l, mid)
	mergeSort(a, temp, mid+1, r)

	i, j, x := l, mid+1, l

	for i <= mid && j <= r {
		for ; i <= mid && a[i] <= a[j]; i++ {
			temp[x] = a[i]
			x++
		}
		for ; j <= r && a[j] <= a[i]; j++ {
			temp[x] = a[j]
			x++
		}
	}
	if i <= mid {
		copy(temp[x:], a[i:mid+1])
		x += mid - i + 1
	}
	if j <= r {
		copy(temp[x:], a[j:r+1])
	}

	copy(a[l:], temp[l:r+1])
}
