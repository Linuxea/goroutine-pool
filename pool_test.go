package goroutinepool

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {

	p := NewPool(10)
	defer p.ShutdownGracefully()

	for i := 0; i < 100; i++ {
		func(ii int) {
			p.Run(func() {
				fmt.Println(ii)
			})
		}(i)
	}

}
