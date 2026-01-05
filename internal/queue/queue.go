package queue

import "Assignment_2/internal/model"

// Simple queue using buffered channel
type TaskQueue struct {
	ch chan model.Task // Buffered channel
}

// Create new queue with specified size
func NewTaskQueue(size int) *TaskQueue {
	return &TaskQueue{
		ch: make(chan model.Task, size), // Buffered channel
	}
}

// Add task to queue
func (q *TaskQueue) Push(task model.Task) {
	q.ch <- task // Send task to channel
}

// Get task from queue (blocks if empty)
func (q *TaskQueue) Pop() model.Task {
	return <-q.ch // Receive task from channel
}

// Close the queue
func (q *TaskQueue) Close() {
	close(q.ch)
}
