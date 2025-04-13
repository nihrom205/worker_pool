package main

import (
	"fmt"
	"time"
)

type Pool struct {
	pool chan struct{}
}

func NewPool(numberWorker int) Pool {
	pool := make(chan struct{}, numberWorker)
	for i := 0; i < numberWorker; i++ {
		pool <- struct{}{}
	}
	return Pool{pool: pool}
}

func (p *Pool) work(n int) {
	<-p.pool
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(n, "pool start")
		p.pool <- struct{}{}
	}()

}

func main() {
	pool := NewPool(3)
	for i := 0; i < 10; i++ {
		pool.work(i)
	}
	time.Sleep(20 * time.Second)
}
