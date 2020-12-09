package main

import "fmt"
// 给定一个按照升序排列的长度为n的整数数组，以及 q 个查询。
//
//对于每个查询，返回一个元素k的起始位置和终止位置（位置从0开始计数）。
//
//如果数组中不存在该元素，则返回“-1 -1”。
//
//输入格式
//第一行包含整数n和q，表示数组长度和询问个数。
//
//第二行包含n个整数（均在1~10000范围内），表示完整数组。
//
//接下来q行，每行包含一个整数k，表示一个询问元素。
//
//输出格式
//共q行，每行包含两个整数，表示所求元素的起始位置和终止位置。
//
//如果数组中不存在该元素，则返回“-1 -1”。
//
//数据范围
//1≤n≤100000
//1≤q≤10000
//1≤k≤10000
//输入样例：
//6 3
//1 2 2 3 3 4
//3
//4
//5
//输出样例：
//3 4
//5 5
//-1 -1
func main() {
	//var n, q int
	//fmt.Scanf("%d %d\n", &n, &q)
	//a := make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Scanf("%d", &a[i])
	//}
	//k := make([]int, q)
	//for i := 0; i < q; i++ {
	//	fmt.Scanf("%d", &k[i])
	//}
	q := 3
	a := []int{1, 2, 3}
	k := []int{1, 2, 3}

	for i := 0; i < q; i++ {
		l, r := findRange(a, k[i])
		fmt.Println(l, r)
	}
}

func findRange(a []int, k int) (int, int) {
	i, j := 0, len(a)-1
	l, r := -1, -1
	mid := 0

	for i < j {
		mid = i + (j-i)>>1

		if a[mid] < k {
			i = mid + 1
		} else {
			j = mid
		}
	}

	if a[i] == k {
		l = i

		i, j = l, len(a)-1

		for i < j {
			mid = i + (j-i+1)>>1

			if a[mid] > k {
				j = mid - 1
			} else {
				i = mid
			}
		}

		if a[j] == k {
			r = j
		}
	}

	return l, r
}
