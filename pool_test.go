package goroutinepool

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {

	p := NewPool(1)
	for i := 0; i < 30; i++ {
		func(ii int) {
			p.Run(func() {
				fmt.Println(ii)
			})
		}(i)
	}

	p.ShutdownGracefully()
}
