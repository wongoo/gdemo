// authors: wangoo
// created: 2018-07-17
// goroutines test

package main

import "fmt"

var limit = make(chan int, 3)
var work = []func(){w1, w2, w3, w4}

func main() {
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select {}
}

func w1() {
	fmt.Println(1)
}
func w2() {
	fmt.Println(2)
}
func w3() {
	fmt.Println(3)
}
func w4() {
	fmt.Println(4)
}
