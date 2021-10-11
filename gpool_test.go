package gpool

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGPool(t *testing.T) {
	pool := New(10)
	fmt.Println("default_pool_size=", runtime.NumGoroutine())
	for i := 0; i < 50; i++ {
		pool.Add(1)
		go func(i int) {
			defer pool.Done()
			fmt.Printf("pool_size=%d | i=%d\n", runtime.NumGoroutine(), i+1)
		}(i)
	}
	pool.Wait()
	fmt.Println("done_pool_size=", runtime.NumGoroutine())
}
