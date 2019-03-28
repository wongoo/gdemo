package chans

import (
	"testing"
)

func BenchmarkChanPoolPerformance(b *testing.B) {
	pools := make(chan int, 1000)

	n := b.N
	go func() {
		for j := 0; j < n; j++ {
			pools <- j
		}
	}()

	for i := 0; i < n; i++ {
		<-pools
	}
}
