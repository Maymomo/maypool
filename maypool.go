package maypool

import (
	"runtime"
)

type Handle func()
type task struct {
	handle Handle
}
type Maypool struct {
	workers  int
	nworker  int
	taskChan chan task
}

func NewPool(workers int) Maypool {
	if workers <= 0 {
		workers = runtime.NumCPU()
	}
	return Maypool{workers: workers, nworker: 0, taskChan: make(chan task, workers*2)}
}

func (pool *Maypool) Process(handle Handle) {
	if pool.nworker < pool.workers {
		pool.newWorker()
	}
	pool.taskChan <- task{handle}
}

func (pool *Maypool) Shutdown() {
	close(pool.taskChan)
}

func (pool *Maypool) newWorker() {
	pool.nworker++
	go func(i int) {
		for {
			task_, ok := <-pool.taskChan
			if !ok {
				break
			}
			task_.handle()
		}
	}(pool.nworker - 1)
}
