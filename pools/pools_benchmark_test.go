package pools

import (
	"bytes"
	"sync"
	"testing"
)

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

func TestPool1(t *testing.T) {
	p1 := sync.Pool{New: func() interface{} {
		return new(bytes.Buffer)
	}}

	b := p1.Get().(*bytes.Buffer)
	b.Write([]byte("hello"))

	p1.Put(b)
}

func TestPool2(t *testing.T) {
	p2 := sync.Pool{New: func() interface{} {
		return make([]byte, 0, 64)
	}}

	b := p2.Get().([]byte)[:0]
	b = append(b, []byte("hello")...)
	t.Logf("addr: %p, val: %s", b, b)
	p2.Put(b)

	b = p2.Get().([]byte)[:0]
	b = append(b, []byte("hello")...)
	t.Logf("addr: %p, val: %s", b, b)
	p2.Put(b)

	b = p2.Get().([]byte)[:0]
	b = append(b, []byte("hello")...)
	t.Logf("addr: %p, val: %s", b, b)
	p2.Put(b)
}
