// authors: wangoo
// created: 2018-05-25
// timer demo

package main

import (
	"time"
	"fmt"
)

func main() {
	tm := time.NewTimer(time.Duration(10 * time.Second))

	fmt.Println(tm)

}
