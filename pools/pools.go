package pools

import (
	"sync/atomic"
)

// 数组池的想法违背了go的思想：Don't communicate by sharing memory; share memory by communicating.
// 而且通过share memory很难沟通，除非goroutine一直运行; 还是用chan沟通比较简单
// 而且通过share memory很难通过单一的变量控制状态
// 性能也并不会很好
type ArrayPool struct {
	r     int32
	w     int32
	count int32
	size  int32
	arr   *[]int
}

func New(arr *[]int) *ArrayPool {
	p := &ArrayPool{
		r:     -1,
		w:     -1,
		count: 0,
		size:  int32(len(*arr)),
		arr:   arr,
	}
	return p
}

func (p *ArrayPool) Get() int {
	if atomic.LoadInt32(&p.count) <= 0 {
		// 应该要等待
		return -1
	}
	idx := atomic.AddInt32(&p.r, 1)
	num := idx % p.size
	if idx >= p.size {
		atomic.CompareAndSwapInt32(&p.r, idx, num)
	}
	atomic.AddInt32(&p.count, -1)
	return (*p.arr)[num]
}

func (p *ArrayPool) Add(i int) bool {
	if atomic.LoadInt32(&p.count) >= p.size {
		// 应该要等待
		return false
	}
	idx := atomic.AddInt32(&p.w, 1)
	num := idx % p.size
	if idx >= p.size {
		atomic.CompareAndSwapInt32(&p.w, idx, num)
	}
	atomic.AddInt32(&p.count, 1)
	(*p.arr)[num] = i
	return true
}
