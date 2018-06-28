// authors: wangoo
// created: 2018-05-25
// timer demo

package basic

import (
	"time"
	"fmt"
	"testing"
)

func TestTimer(t *testing.T) {
	tm := time.NewTimer(time.Duration(10 * time.Second))

	fmt.Println(tm)

}
