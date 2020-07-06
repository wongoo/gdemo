package main

import (
	"net/http"
	"sync"
)

func newBytePool() sync.Pool {
	return sync.Pool{New: func() interface{} {
		return make([]byte, 0, 512)
	}}
}

var bytePool = newBytePool()

func get() []byte {
	return bytePool.Get().([]byte)
}

func put(b []byte) {
	bytePool.Put(b)
}

func main() {
	b:=get()
	http.Server{}
}
