// authors: wangoo
// created: 2018-05-24
// string demo

package main

import (
	"fmt"
	"time"
	"bytes"
	"strings"
)

func main() {
	// alpha char string
	s := "hello world"
	fmt.Printf("s=%v\n", s)

	// string length
	fmt.Printf("len(s)=%v\n", len(s))

	subSlice := s[3:8]
	fmt.Printf("subSlice=%v\n", subSlice)

	// chinese utf8 string
	cn := "中文字符串你好"
	fmt.Printf("cn=%v\n", cn)

	// len of bytes
	fmt.Printf("len(cn)=%v\n", len(cn))

	// slice sub-string operation on byte, which will get messy sub-string for utf8
	cnSlice := cn[1:3]
	fmt.Printf("cnSlice=%v\n", cnSlice)

	// string to rune
	cnRune := []rune(cn)
	fmt.Printf("cnRune=%v\n", cnRune)

	cnRuneStr := string(cnRune)
	fmt.Printf("cnRuneStr=%v\n", cnRuneStr)

	cnRuneSub := string(cnRune[2:5])
	fmt.Printf("cnRuneSub=%v\n", cnRuneSub)

	//string contract
	s1 := "part1"
	s2 := "part2"
	s3 := s1 + s2
	fmt.Printf("s1=%v\n", s1)
	fmt.Printf("s2=%v\n", s2)
	fmt.Printf("s3=%v\n", s3)

	stringContractPerformance()
}

func stringContractPerformance() {
	loopTime := 10000
	//第一种连接方法（最快）
	var buffer bytes.Buffer
	s := time.Now()
	for i := 0; i < loopTime; i++ {
		buffer.WriteString("test is here\n")
	}
	buffer.String() // 拼接结果
	e := time.Now()
	fmt.Println("1 time is ", e.Sub(s).Seconds())

	//第二种方法
	s = time.Now()
	var sl []string
	for i := 0; i < loopTime; i++ {
		sl = append(sl, "test is here\n")
	}
	strings.Join(sl, "")
	e = time.Now()
	fmt.Println("2 time is", e.Sub(s).Seconds())

	//第三种方法
	s = time.Now()
	str := ""
	for i := 0; i < loopTime; i++ {
		str += "test is here\n"
	}
	e = time.Now()
	fmt.Println("3 time is ", e.Sub(s).Seconds())

	//第四种方法
	s = time.Now()
	str4 := ""
	for i := 0; i < loopTime; i++ {
		str4 = str4 + "test is here"
	}
	e = time.Now()
	fmt.Println("4 time is ", e.Sub(s).Seconds())
}
