package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"
	"testing"
	"time"
)

type task func()

func doFunc(f task) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "%s goroutine panic: %v\n%s\n",
				time.Now(), r, string(debug.Stack()))
		}
	}()
	f()
}

var ff = func() {
	fmt.Fprintln(ioutil.Discard, "hello world")
}

func BenchmarkAAA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Fprintf(os.Stderr, "%s goroutine panic: %v\n%s\n",
						time.Now(), r, string(debug.Stack()))
				}
			}()
			ff()
		}()
	}
}

func BenchmarkBBB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doFunc(ff)
	}
}
