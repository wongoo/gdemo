package atoms

import (
	"sync/atomic"
	"testing"
)

func BenchmarkRemainder(b *testing.B) {
	var num int32
	var remainderBase int32 = 1234
	for i := 0; i < b.N; i++ {
		add := atomic.AddInt32(&num, 1)
		_ = add % remainderBase
	}
}

func BenchmarkRemainder2(b *testing.B) {
	var num int32
	var remainderBase int32 = 1234
	for i := 0; i < b.N; i++ {
		add := atomic.AddInt32(&num, 1)
		if add > remainderBase {
			r := add % remainderBase
			// try to reset the num
			atomic.CompareAndSwapInt32(&num, add, r)
		}
	}
}
