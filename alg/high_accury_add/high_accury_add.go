package main

import (
	"fmt"
)

// 给定两个正整数，计算它们的和。
//
//输入格式
//共两行，每行包含一个整数。
//
//输出格式
//共一行，包含所求的和。
//
//数据范围
//1≤整数长度≤100000
//输入样例：
//12
//23
//输出样例：
//35
func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	result := make([]int, 100010)

	i := len(a) - 1
	j := len(b) - 1
	n := 0
	t := 0

	for i >= 0 || j >= 0 {
		if i >= 0 {
			t += int(a[i] - '0')
		}
		if j >= 0 {
			t += int(b[j] - '0')
		}

		result[n] = t % 10

		t = t / 10

		n++
		i--
		j--
	}

	if t > 0 {
		fmt.Print(t)
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
}
