package goroutinepool

import (
	"sync"
)

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func NewPool(maxGoroutine uint) *Pool {
	pool := &Pool{
		work: make(chan Worker, maxGoroutine),
	}
	pool.wg.Add(int(maxGoroutine))
	for i := 0; i < int(maxGoroutine); i++ {
		go func() {
			for {
				w, ok := <-pool.work
				if !ok {
					break
				}
				w.work()
			}
			pool.wg.Done()
		}()

	}

	return pool
}

func (p *Pool) Run(fun func()) {
	p.work <- WorkFun(fun)
}

func (p *Pool) ShutdownGracefully() {
	close(p.work)
	p.wg.Wait()
}
