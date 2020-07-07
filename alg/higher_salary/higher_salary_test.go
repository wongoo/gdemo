package higher_salary

import "testing"

//示例1：
//输入：
//3 3
//1 100
//10 1000
//1000000000 1001
//9 10 1000000000
//输出：
//100
//1000
//1001
func TestHigherSalary(t *testing.T) {
	HigherSalary([]int{1, 10, 1000000000}, []int{100, 1000, 1001}, []int{9, 10, 1000000000})
}
