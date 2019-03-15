// authors: wangoo
// created: 2018-05-24
// array demo

package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	// define array type and size
	var arr [10]int
	arr[0] = 43
	arr[1] = 12

	fmt.Printf("arr[0]=%d\n", arr[0])

	// define and initial array
	arr1 := [3]int{1, 2, 3}
	fmt.Printf("arr1[1]=%d\n", arr1[1])

	// ignore array length , calculate array length automatically
	arr2 := [...]int{4, 5, 6}
	fmt.Printf("arr2[1]=%d\n", arr2[1])

	// two-dimensional array ;
	doubleArr := [2][4]int{{1, 2, 3, 4}, {6, 7, 8, 9}}
	fmt.Printf("doubleArr[1][3]=%d\n", doubleArr[1][3])

	// slice 与 array 接近，但是在新的元素加入的时候可以增加长度。slice 总是指向底层的一个array。
	// slice 是一个指向 array 的指针，这是其与 array 不同的地方;
	// slice 是引用类型， 这意味着当赋值某个 slice 到另外一个变量，两个引用会指向同一个array.
	arr3 := [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	// 给定一个 array 或者其他 slice，一个新 slice 通过 a[I:J] 的方式创建。
	// 这会创建一个新 的 slice，指向变量 a，从序号 I 开始，结束在序号 J之前。长度为 J - I。
	// array[n:m] 从 array 创建了一个 slice，具有元素 n 到 m-1
	fmt.Printf("arr3=%v\n", arr3)
	fmt.Printf("arr3[3:]=%v\n", arr3[3:])
	fmt.Printf("arr3[:3]=%v\n", arr3[:3])

	fmt.Printf("arr4=%v\n", arr4)
	fmt.Printf("arr4[4:]=%v\n", arr4[4:])
	fmt.Printf("arr4[2:4]=%v\n", arr4[2:4])

	// 函数 append 向 slice s 追加零值或其他 x 值，并且返回追加后的新的、与 s 有相同类型的 slice。
	// 如果 s 没有足够的容量存储追加的值，append 分配一 个足够大的、新的 slice 来存放原有 slice 的元素和追加的值。
	// 因此，返回 的 slice 可能指向不同的底层 array。
	sliceA := arr3[:3]
	fmt.Printf("sliceA=%v\n", sliceA)

	// slice append
	sliceB := append(sliceA, 'x')
	sliceC := append(sliceB, 'y')
	fmt.Printf("[after appending] sliceA=%v\n", sliceA)
	fmt.Printf("[after appending] sliceB=%v\n", sliceB)
	fmt.Printf("[after appending] sliceC=%v\n", sliceC)

	// update slice will update referencing array
	sliceA[0] = '0'
	sliceB[1] = '1'
	sliceC[3] = '2'
	fmt.Printf("[after update] sliceA=%v\n", sliceA)
	fmt.Printf("[after update] sliceB=%v\n", sliceB)
	fmt.Printf("[after update] sliceC=%v\n", sliceC)

}

func TestArrayAddr(t *testing.T) {
	arr := []int{0}
	for i := 0; i < 10; i++ {
		fmt.Printf("arr addr: %p\n", arr)
		arr = append(arr, i)
	}

	sliceType := reflect.SliceOf(reflect.TypeOf(1))
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	for i := 0; i < 10; i++ {
		fmt.Printf("slice value addr: %d\n", sliceValue.Pointer())
		sliceValue = reflect.Append(sliceValue, reflect.ValueOf(i))
	}

}
