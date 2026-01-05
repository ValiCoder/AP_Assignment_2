package store

import (
	"Assignment_2/internal/model"
	"sync"
)

//I can/t lie, I didn't do it all by myself. I had deepseek and chatgpt help me, tell me what to do.
//I'm sorry

type TaskRepository struct {
	mu    sync.RWMutex
	tasks map[string]model.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]model.Task),
	}
}

func (r *TaskRepository) Add(task model.Task) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
}

func (r *TaskRepository) Get(id string) (model.Task, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, exists := r.tasks[id]
	return task, exists
}

func (r *TaskRepository) GetAll() []model.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]model.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		result = append(result, task)
	}
	return result
}

func (r *TaskRepository) Update(task model.Task) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return false
	}

	r.tasks[task.ID] = task
	return true
}

func (r *TaskRepository) GetStats() map[string]int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	stats := map[string]int{
		"PENDING":     0,
		"IN_PROGRESS": 0,
		"DONE":        0,
	}

	for _, task := range r.tasks {
		stats[task.Status]++
	}

	return stats
}
