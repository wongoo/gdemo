package defergo

import (
	"sync"
	"testing"
)

var (
	lockDefer   sync.Mutex
	lockNoDefer sync.Mutex
	seq         int
)

func addDefer() {
	lockDefer.Lock()
	defer lockDefer.Unlock()
	seq++
}

func addNoDefer() {
	lockNoDefer.Lock()
	seq++
	lockNoDefer.Unlock()
}

func BenchmarkDeferPerf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addDefer()
	}
}

func BenchmarkNoDeferPerf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addNoDefer()
	}
}
