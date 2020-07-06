package chans

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestOneBufferChannel(t *testing.T) {
	c := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		log.Println("receive 1 start")
		<-c
		log.Println("receive 1 end")

		time.Sleep(time.Second * 2)

		log.Println("receive 2 start")
		<-c
		log.Println("receive 2 end")

		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 2)

		log.Println("send 1")
		c <- struct{}{}
		log.Println("send 1 finish")

		log.Println("send 2")
		c <- struct{}{}
		log.Println("send 2 finish")

		wg.Done()
	}()

	wg.Wait()
}
