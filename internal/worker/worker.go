package worker

import (
	"Assignment_2/internal/model"
	"Assignment_2/internal/store"
	"log"
	"time"
)

type Worker struct {
	ID     int
	taskCh <-chan model.Task // Read-only channel
	repo   *store.Repository
	stopCh chan struct{} // Signal to stop
}

func NewWorker(id int, taskCh <-chan model.Task, repo *store.Repository) *Worker {
	return &Worker{
		ID:     id,
		taskCh: taskCh,
		repo:   repo,
		stopCh: make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go w.run() // Start worker in goroutine
}

func (w *Worker) Stop() {
	close(w.stopCh)
}

func (w *Worker) run() {
	log.Printf("Worker %d started", w.ID)

	for {
		select {
		case <-w.stopCh:
			log.Printf("Worker %d stopping", w.ID)
			return

		case task := <-w.taskCh:
			w.processTask(task)
		}
	}
}

func (w *Worker) processTask(task model.Task) {
	log.Printf("Worker %d processing task %s", w.ID, task.ID)

	// 1. Update status to IN_PROGRESS
	task.Status = "IN_PROGRESS"
	w.repo.Update(task)

	// 2. Simulate work (2 seconds)
	time.Sleep(2 * time.Second)

	// 3. Update status to DONE
	task.Status = "DONE"
	w.repo.Update(task)

	log.Printf("Worker %d completed task %s", w.ID, task.ID)
}
