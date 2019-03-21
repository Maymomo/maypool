package maypool

import (
	"sync"
	"testing"
	"time"
)

func TestMaypool(t *testing.T) {
	var mutex sync.Mutex
	pool := NewPool(10)
	var i int = 0
	c := make(chan int, 10000)
	for ii := 0; ii < 10000; ii++ {
		pool.Process(func() {
			mutex.Lock()
			i++
			mutex.Unlock()
			c <- 1
		})
	}

	for ii := 0; ii < 10000; ii++ {
		<-c
	}
	if i != 10000 {
		t.Fatal("i do not equal 10000")
	}
	pool.Shutdown()
	time.Sleep(5 * time.Second)
}
