package api

import (
	"Assignment_2/internal/model"
	"Assignment_2/internal/store"
	"Assignment_2/internal/worker"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	repo       *store.Repository
	workerPool *worker.WorkerPool // Add worker pool
}

func NewHandler(repo *store.Repository, workerPool *worker.WorkerPool) *Handler {
	return &Handler{
		repo:       repo,
		workerPool: workerPool,
	}
}

// POST /tasks
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Payload string `json:"payload"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Create task
	task := model.Task{
		ID:      uuid.New().String(),
		Payload: req.Payload,
		Status:  "PENDING",
	}

	// Save to repository
	h.repo.Add(task)

	// Submit to worker pool for processing
	h.workerPool.SubmitTask(task)

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"id": task.ID,
	})
}

// GET /tasks
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks := h.repo.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GET /tasks/{id}
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get ID from URL path
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	task, exists := h.repo.Get(id)
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// GET /stats
func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stats := h.repo.GetStats()

	// Format as required by assignment
	response := map[string]int{
		"submitted":   stats["PENDING"],
		"completed":   stats["DONE"],
		"in_progress": stats["IN_PROGRESS"],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
