[![Go Report Card](https://goreportcard.com/badge/github.com/umitanilkilic/gorev)](https://goreportcard.com/report/github.com/umitanilkilic/gorev)
 # Gorev
The `gorev` package provides a simple and efficient task management system in Go. It includes a worker structure for managing tasks with priorities and performing them asynchronously. This package is designed to be easy to use and customizable for various applications requiring task execution.
## Overview

This package provides a worker implementation for managing and executing tasks concurrently in Go.

## Key Features

- Task Management:
    - Adds tasks to a worker's queue.
    - Sorts tasks based on priority.
    - Removes tasks by index.
- Task Execution:
    - Continuously performs tasks in the queue.
    - Handles errors during task execution.
- Worker State:
    - Starts and stops the worker.
    - Checks if the worker is running.

## Usage

1. Create a Worker:
   ```go
   worker := NewWorker()
   ```
2. Create Tasks:
   ```go
   task1 := NewTask(myTaskFunction, 10) // Priority 10
   task2 := NewTask(anotherTaskFunction, 5) // Priority 5
   ```
3. Add Tasks to Worker:
   ```go
   worker.AddTask(task1)
   worker.AddTask(task2)
   ```
4. Start the Worker:
   ```go
   worker.Start()
   ```
5. Stop the Worker (when needed):
   ```go
   worker.Stop()
   ```

## Notes

- **Priority-Based Task Execution:** Tasks are executed in descending order of priority.
- **Concurrency:** Tasks are executed concurrently using goroutines.
- **Error Handling:** Errors during task execution are logged and the worker continues with other tasks.

**## TODOs**

- **Improve Task Sorting Efficiency:** Explore more efficient sorting algorithms for large task queues.
- **Implement Task Timeouts:** Add a mechanism to handle tasks that exceed a specified timeout.
