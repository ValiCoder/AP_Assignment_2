package worker

import (
	"Assignment_2/internal/model"
	"Assignment_2/internal/store"
)

type WorkerPool struct {
	workers []*Worker
	queue   *TaskQueue // We'll create this type
}

// Simple wrapper for queue (temporary)
type TaskQueue struct {
	ch chan model.Task
}

func NewWorkerPool(numWorkers int, queueSize int, repo *store.Repository) *WorkerPool {
	// Create task queue
	taskCh := make(chan model.Task, queueSize)

	// Create workers
	workers := make([]*Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = NewWorker(i+1, taskCh, repo)
	}

	return &WorkerPool{
		workers: workers,
		queue:   &TaskQueue{ch: taskCh},
	}
}

func (wp *WorkerPool) Start() {
	for _, worker := range wp.workers {
		worker.Start()
	}
}

func (wp *WorkerPool) Stop() {
	for _, worker := range wp.workers {
		worker.Stop()
	}
	close(wp.queue.ch)
}

func (wp *WorkerPool) SubmitTask(task model.Task) {
	wp.queue.ch <- task
}
