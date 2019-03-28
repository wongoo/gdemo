package pools

import "testing"

func BenchmarkArrayPoolPerformance(b *testing.B) {
	arr := make([]int, 1000, 1000)
	p := New(&arr)

	n := b.N
	go func() {
		for j := 0; j < n; j++ {
			p.Add(j)
		}
	}()

	for i := 0; i < n; i++ {
		p.Get()
	}
}
