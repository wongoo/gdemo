package main

import "fmt"

func cmp(a, b string) int {
	if len(a) != len(b) {
		if len(a) > len(b) {
			return 1
		} else {
			return -1
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if a[i] > b[i] {
				return 1
			} else {
				return -1
			}
		}
	}
	return 0
}

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	negative := false
	switch cmp(a, b) {
	case 0:
		fmt.Print(0)
		return
	case -1:
		a, b = b, a
		negative = true
	}

	r := make([]int, len(a)+10)
	n := 0
	t := 0

	for i, j := len(a)-1, len(b)-1; i >= 0; i-- {
		t = int(a[i]-'0') - t
		if j >= 0 {
			t -= int(b[j] - '0')
			j--
		}

		r[n] = (t + 10) % 10
		n++

		if t < 0 {
			t = 1
		} else {
			t = 0
		}
	}

	if negative {
		fmt.Print("-")
	}

	started := false
	for i := n - 1; i >= 0; i-- {
		if started || r[i] > 0 {
			fmt.Print(r[i])
			started = true
		}
	}
}
