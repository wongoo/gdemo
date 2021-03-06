// authors: wangoo
// created: 2018-05-24
// map demo

// map是无序的
// map的长度是不固定的

// var numberMap map[string] int  // 两种申明方式
// make用于内建类型(map、slice和channel)的内存分配。new用于各种类型的内存分配。
// make只能创建slice、map和channel，并且返回一个有初始值(非零)的T 类型 make返回初始化后的(非零)值。

//各类型初始值
//	int		0
//	int8	0
//	int32	0
//	int64	0
//	uint	0x0
//	rune	0
//	byte	0x0
//	float32	0
//	float64	0
//	boolean	false
//	string	""

package basic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	numberMap := make(map[string]int)

	// check
	fmt.Printf("init --> numberMap[\"one\"]=%d\n", numberMap["one"])

	// assign
	numberMap["one"] = 1
	fmt.Printf("assign --> numberMap[\"one\"]=%d\n", numberMap["one"])

	// update
	numberMap["one"] = 11
	fmt.Printf("updated --> numberMap[\"one\"]=%d\n", numberMap["one"])
	_, ok := numberMap["one"]
	fmt.Printf("contains --> numberMap[\"one\"]=%v\n", ok)

	// deleted
	delete(numberMap, "one")
	fmt.Printf("deleted --> numberMap[\"one\"]=%d\n", numberMap["one"])
	_, ok = numberMap["one"]
	fmt.Printf("contains --> numberMap[\"one\"]=%v\n", ok)

	// init and assign
	numberMap = map[string]int{"c": 4, "d": 23}
	fmt.Printf("len(numberMap)=%d\n", len(numberMap))
}

func TestMapAddress(t *testing.T) {
	m := map[int]int{}
	for i := 0; i < 20; i++ {
		fmt.Printf("map addr: %p\n", m)
		m[i] = i
	}
	typ := reflect.TypeOf(1)
	mapTyp := reflect.MapOf(typ, typ)
	mapValue := reflect.MakeMap(mapTyp)

	for i := 0; i < 20; i++ {
		fmt.Printf("map value addr: %d\n", mapValue.Pointer())
		iv := reflect.ValueOf(i)
		mapValue.SetMapIndex(iv, iv)
	}
}

func TestSyncMap(t *testing.T) {
	//m:=sync.Map{}
}

func TestMapFuncRef(t *testing.T) {
	m := make(map[string]string)

	addMapItem(m)
	fmt.Printf("map address: %p\n", m)

	if _, ok := m["a"]; !ok {
		t.Error("should contains key add in calling func")
	}

	assert.Equal(t, 1, len(m))
}

func addMapItem(m map[string]string) {
	m["a"] = "b"
	fmt.Printf("map address: %p\n", m)
}
