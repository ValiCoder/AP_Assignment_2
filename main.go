package main

import (
	"Assignment_2/internal/api"
	"Assignment_2/internal/store"
	"Assignment_2/internal/worker"
	"fmt"
	"net/http"
)

func testGenerics() {
	fmt.Println("=== Testing Generic Repository ===")

	// Test 1: String repository
	stringRepo := store.NewGenericRepository[string, string]()
	stringRepo.Create("key1", "value1")
	stringRepo.Create("key2", "value2")

	if val, exists := stringRepo.Get("key1"); exists {
		fmt.Printf("String repo: key1 = %s\n", val)
	}

	// Test 2: Int repository
	intRepo := store.NewGenericRepository[string, int]()
	intRepo.Create("age", 25)
	intRepo.Create("count", 100)

	if val, exists := intRepo.Get("age"); exists {
		fmt.Printf("Int repo: age = %d\n", val)
	}

	// Test 3: Get all
	allStrings := stringRepo.GetAll()
	fmt.Printf("All strings: %v\n", allStrings)
}

func main() {

	// 1. Create repository
	repo := store.NewRepository()

	// 2. Create worker pool (2 workers, queue size 100)
	workerPool := worker.NewWorkerPool(2, 100, repo)

	// 3. Start workers
	workerPool.Start()
	defer workerPool.Stop() // Will stop when main exits

	// 4. Create HTTP handler
	handler := api.NewHandler(repo, workerPool)

	// 5. Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks", handler.CreateTask)
	mux.HandleFunc("GET /tasks", handler.GetTasks)
	mux.HandleFunc("GET /tasks/{id}", handler.GetTask)
	mux.HandleFunc("GET /stats", handler.GetStats)

	// 6. Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Worker pool started with 2 workers")
	fmt.Println("Test endpoints:")
	fmt.Println("  POST   /tasks")
	fmt.Println("  GET    /tasks")
	fmt.Println("  GET    /tasks/{id}")
	fmt.Println("  GET    /stats")

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
