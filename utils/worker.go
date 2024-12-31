package utils

import (
	"sync"
)

type WorkerPool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

func NewWorkerPool(poolSize int) *WorkerPool {
	wp := &WorkerPool{
		tasks: make(chan func()),
	}

	for i := 0; i < poolSize; i++ {
		go wp.worker()
	}

	return wp
}

func (wp *WorkerPool) AddTask(task func()) {
	wp.wg.Add(1)
	wp.tasks <- task
}

func (wp *WorkerPool) worker() {
	for task := range wp.tasks {
		task()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Wait() {
	close(wp.tasks)
	wp.wg.Wait()
}
