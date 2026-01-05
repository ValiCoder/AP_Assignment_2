# Task Manager Service

A concurrent HTTP service for managing background tasks with asynchronous processing.

## Features

- REST API for task management
- Concurrent processing with worker pool  
- Background monitoring of task statistics
- Graceful shutdown handling
- Thread-safe shared state management
- Generic in-memory repository

## API Endpoints

### POST /tasks
Create a new task.

Request:
```json
{
  "payload": "task content"
}
```

Response (201 Created):
```json
{
  "id": "unique-task-id"
}
```

### GET /tasks
Get all tasks.

Response:
```json
[
  {
    "id": "task-id",
    "payload": "task content",
    "status": "PENDING"
  }
]
```

### GET /tasks/{id}
Get a specific task by ID.

Response (200 OK):
```json
{
  "id": "task-id",
  "payload": "task content", 
  "status": "DONE"
}
```

Response (404 Not Found): If task doesn't exist.

### GET /stats
Get server statistics.

Response:
```json
{
  "submitted": 5,
  "completed": 3,
  "in_progress": 1
}
```

## Architecture

- Task Repository: Generic in-memory storage with thread-safe access
- Worker Pool: 2+ concurrent workers processing tasks asynchronously
- Task Queue: Buffered channel for task distribution
- Background Monitor: Logs task statistics every 5 seconds
- Graceful Shutdown: Proper cleanup on termination signal

## Requirements

### Part A: Task API (40%)
- All required endpoints implemented
- JSON request/response handling
- Proper HTTP status codes

### Part B: Concurrency & Worker Pool (25%)
- Buffered channel task queue
- 2+ worker goroutines
- Thread-safe state updates with mutexes

### Part C: Background Monitoring (15%)
- Periodic logging every 5 seconds
- Uses time.Ticker
- Controlled with stop channel

### Part D: Graceful Shutdown (10%)
- OS signal capture (Ctrl+C)
- Stop accepting new requests
- Stop workers and monitor
- Allow active tasks to complete

### Part E: Generics (10%)
- Generic repository implementation
- Type-safe storage for any key-value types

## Running the Service

1. Install dependencies:
```
go mod tidy
```

2. Run the server:
```
go run main.go
```

3. Server starts on: http://localhost:8080

## Testing

Create a task:
```
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"payload": "test task"}'
```

Get all tasks:
```
curl http://localhost:8080/tasks
```

Get statistics:
```
curl http://localhost:8080/stats
```

## Project Structure

```
Assignment_2/
-- go.mod
-- main.go
-- README.md
-- internal/
   -- model/       # Data structures
   -- store/       # Generic repository
   -- api/         # HTTP handlers
   -- worker/      # Worker pool
   -- queue/       # Task queue
```

## Implementation Details

- Concurrency: Uses Go channels and goroutines for async processing
- Synchronization: sync.RWMutex for thread-safe data access
- Error Handling: Proper HTTP error responses
- Code Organization: Clean separation of concerns

## Shutdown Procedure

Press Ctrl+C to initiate graceful shutdown:
1. Stop accepting new HTTP requests
2. Stop background monitor
3. Stop worker pool
4. Wait for active tasks to complete
5. Exit cleanly

## Notes

- All data stored in memory (no persistence)
- Task processing simulated with 2-second delays
- Unique task IDs generated using UUID
- Monitor logs to stdout every 5 seconds
