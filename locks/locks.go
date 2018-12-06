// authors: wangoo
// created: 2018-07-17
// locks

package main

import (
	"sync"
	"fmt"
	"time"
)

var l sync.Mutex
var wr sync.RWMutex
var a string
var c int

func f() {
	a = "hello, world"
	l.Unlock()
}

func fr() {
	wr.RLock()
	fmt.Println("read", c)
	time.Sleep(time.Second)
	wr.RUnlock()
}

func fw() {
	wr.Lock()

	c++
	fmt.Println("write", c)
	time.Sleep(time.Second)
	wr.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
	l.Unlock()

	for i := 0; i < 10; i++ {
		go fw()
		go fr()
	}

	select {}
}
