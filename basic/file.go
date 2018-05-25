// authors: wangoo
// created: 2018-05-25
// file demo

package main

import (
	"os"
	"fmt"
)

func main() {
	path := "test.txt"
	file, _ := os.Create(path)

	fmt.Println("write file")
	file.WriteString("hello")

	fmt.Println("sync from memory to disk")
	file.Sync()

	fmt.Println("close file")
	file.Close()

	file, ok := os.Open(path)
	if ok == nil {
		fmt.Printf("file %v exists\n", path)
	} else {
		fmt.Printf("file %v not exists\n", path)
	}
	b := make([]byte, 10)

	n, _ := file.Read(b)
	fmt.Printf("read %v bytes\n", n)
	fmt.Printf("read bytes:%v\n", string(b))


	fmt.Println("remove file")
	os.Remove(path)

	_, ok = os.Open(path)
	if ok == nil {
		fmt.Printf("file %v exists\n", path)
	} else {
		fmt.Printf("file %v not exists\n", path)
	}

}
