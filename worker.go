package gorev

import (
	"errors"
	"fmt"
	"sort"

	"github.com/google/uuid"
)

type Worker struct {
	WorkerId  uint32
	tasks     []*Task
	IsRunning bool
	stopChan  chan bool
}

func NewWorker() *Worker {
	return &Worker{WorkerId: uuid.New().ID(), stopChan: make(chan bool)}
}

func (w *Worker) AddTask(task *Task) {
	//Add task to worker
	w.tasks = append(w.tasks, task)
	//Sort tasks
	w.sortTasks()
}

func (w *Worker) sortTasks() {
	if len(w.tasks) <= 1 {
		return
	}
	///TODO: Bunu değiştir çünkü çok zaman alıyor mesela önceliği ilk elemandan yüksek ise en başa koy vs vs sıralama algoritmalarını araştır.
	sort.Slice(w.tasks, func(i, j int) bool {
		return w.tasks[i].Priority > w.tasks[j].Priority
	})
}

func (w *Worker) RemoveTaskByIndex(taskIndex int) error {
	//Check if taskIndex is valid
	if taskIndex < 0 || taskIndex > len(w.tasks) {
		return errors.New("invalid task index")
	}
	//Remove task from worker
	w.tasks = append(w.tasks[:taskIndex], w.tasks[taskIndex+1:]...)

	return nil
}

func (w *Worker) performTasks() error {
	var err error
	for {
		select {
		case <-w.stopChan:
			return nil
		default:
			for _, t := range w.tasks {
				err = t.TaskInterface.Perform()
				if err != nil {
					return fmt.Errorf("error while performing task: %v", err)
				}
			}
			return nil
		}
	}
}

func (w *Worker) GetTasks() []*Task {
	return w.tasks
}

func (w *Worker) Start() error {
	if w.IsRunning {
		return errors.New("worker is already running")
	}

	go w.performTasks()

	w.IsRunning = true
	return nil
}

func (w *Worker) Stop() error {
	if !w.IsRunning {
		return errors.New("worker is not running")
	}
	w.stopChan <- true
	w.IsRunning = false

	return nil
}
