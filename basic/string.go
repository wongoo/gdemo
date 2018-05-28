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
	stringBasic()
	//subString()
	//stringContract()
}

func stringBasic() {
	fmt.Println("-----------------------")
	s := "hello world to every"
	arr := strings.Split(s, " ")
	for n, item := range arr {
		fmt.Println(n, item)
	}
	fmt.Println(strings.Fields(s))
	fmt.Println(strings.Fields("hello  world		hi	 every"))

	fmt.Println(strings.FieldsFunc("hello everyone how are you", func(r rune) bool {
		return r == 'e'
	}))

	fmt.Println(strings.Contains(s, "world"))
	fmt.Println(strings.Contains(s, "worle"))

	fmt.Println(strings.ContainsAny(s, "abc"))
	fmt.Println(strings.ContainsAny(s, "def"))

	fmt.Println(strings.ContainsRune(s, '好'))
	fmt.Println(strings.ContainsRune("你好吗", '好'))

	fmt.Println(strings.Count(s, "o"))
	fmt.Println(strings.Compare("abc", "abc"))
	fmt.Println(strings.Compare("abc", "abcd"))

	fmt.Println(strings.EqualFold("abc", "abc"))
	fmt.Println(strings.EqualFold("你好", "你好"))

	fmt.Println(strings.HasPrefix("abcdefg", "abc"))
	fmt.Println(strings.HasSuffix("abcdefg", "efg"))

	fmt.Println(strings.Index(s, "lo"))
	fmt.Println(strings.IndexRune(s, 'o'))
	fmt.Println(strings.IndexByte(s, 'o'))
	fmt.Println(strings.IndexFunc(s, func(r rune) bool {
		return r == 'o'
	}))

	fmt.Println(strings.LastIndex(s, "o"))
	fmt.Println(strings.LastIndexAny(s, "o"))
	fmt.Println(strings.LastIndexByte(s, 'o'))
	fmt.Println(strings.LastIndexFunc(s, func(r rune) bool {
		return r == 'o'
	}))

	fmt.Println(strings.Map(func(r rune) rune {
		if r == 'o' {
			return -1
		}
		return r
	}, s))

	fmt.Println(strings.Repeat("hello", 3))
	fmt.Println(strings.Replace(s, "o", "0", -1))

	fmt.Println(strings.Title(s))
	fmt.Println(strings.ToLower(s))
	// fmt.Println(strings.ToLowerSpecial(f,s))

	fmt.Println(strings.ToTitle(s))
	fmt.Println(strings.ToUpper(s))
}

func subString() {
	fmt.Println("-----------------------")
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

}

func stringContract() {
	fmt.Println("-----------------------")
	//string contract
	s1 := "part1"
	s2 := "part2"
	s3 := s1 + s2
	fmt.Printf("s1=%v\n", s1)
	fmt.Printf("s2=%v\n", s2)
	fmt.Printf("s3=%v\n", s3)

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
