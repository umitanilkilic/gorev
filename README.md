[![Go Report Card](https://goreportcard.com/badge/github.com/umitanilkilic/gorev)](https://goreportcard.com/report/github.com/umitanilkilic/gorev) [![Go Reference](https://pkg.go.dev/badge/github.com/umitanilkilic/gorev.svg)](https://pkg.go.dev/github.com/umitanilkilic/gorev)
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
   ```go
   type ExampleTaskStruct1 struct{}

func (t *ExampleTaskStruct1) Perform() error {
	return nil
}

type ExampleTaskStruct2 struct{}

func (t *ExampleTaskStruct2) Perform() error {
	return errors.New("error ;(")
}

func main() {
	// Create a new tasks
	task1, err := gorev.NewTask(&ExampleTaskStruct1{}, 3)
	if err != nil {
		fmt.Println(err)
	}

	task2, err := gorev.NewTask(&ExampleTaskStruct2{}, 9)
	if err != nil {
		fmt.Println(err)
	}

	// Create a new worker
	worker := gorev.NewWorker()

	// Add tasks to the worker
	worker.AddTask(task1)
	worker.AddTask(task2)

	// Start the worker
	worker.Start()

	// Error Report
	fmt.Printf("%v", <-worker.GetErrorReports())

	// Stop the worker
	worker.Stop()
}
   ```

## Notes

- **Priority-Based Task Execution:** Tasks are executed in descending order of priority.
- **Concurrency:** Tasks are executed concurrently using goroutines.
- **Error Reporting**

## TODOs

- **Improve Task Sorting Efficiency:** Explore more efficient sorting algorithms for large task queues.
- **Implement Task Timeouts:** Add a mechanism to handle tasks that exceed a specified timeout.
