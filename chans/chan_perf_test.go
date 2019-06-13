package chans

import (
	"testing"
)

func BenchmarkChanPoolPerformance(b *testing.B) {
	pools := make(chan int, 1024)

	go func() {
		for j := 0; j < b.N; j++ {
			go func() {
				for n := 0; n < 100; n++ {
					pools <- j
				}
			}()
		}
	}()

	for i := 0; i < 100*b.N; i++ {
		<-pools
	}
}
