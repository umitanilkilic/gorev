package gorev

import (
	"errors"
	"sort"
	"time"

	"github.com/google/uuid"
)

type Worker struct {
	workerId  uint32
	tasks     []*Task
	isRunning bool
	stopChan  chan bool

	errorReports chan ErrorReport
}
type ErrorReport struct {
	TaskID    uint32
	Error     error
	TimeStamp int64
}

func NewWorker() *Worker {
	return &Worker{workerId: uuid.New().ID(), stopChan: make(chan bool, 1), errorReports: make(chan ErrorReport)}
}

func (w *Worker) GetWorkerId() uint32 {
	return w.workerId
}

func (w *Worker) IsWorkerRunning() bool {
	return w.isRunning
}

func (w *Worker) GetErrorReports() <-chan ErrorReport {
	return w.errorReports
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
	if taskIndex < 0 || taskIndex >= len(w.tasks) {
		return errors.New("invalid task index")
	}
	//Remove task from worker
	w.tasks = append(w.tasks[:taskIndex], w.tasks[taskIndex+1:]...)

	return nil
}

func (w *Worker) RemoveTaskByTaskID(taskID uint32) error {
	//Find task index by taskID
	taskIndex := -1
	for i, t := range w.tasks {
		if t.taskID == taskID {
			taskIndex = i
			break
		}
	}
	//Check if taskIndex is valid
	if taskIndex == -1 {
		return errors.New("task not found")
	}
	//Remove task from worker
	w.tasks = append(w.tasks[:taskIndex], w.tasks[taskIndex+1:]...)

	return nil
}

func (w *Worker) RemoveAllTasks() {
	w.tasks = nil
}

func (w *Worker) performTasks() {
	for {
		select {
		case <-w.stopChan:
			return
		default:
			for k, t := range w.tasks {
				err := t.TaskInterface.Perform()
				if err != nil {
					w.errorReports <- ErrorReport{TaskID: t.taskID, Error: err, TimeStamp: time.Now().Unix()}
				}
				w.RemoveTaskByIndex(k)
			}
			w.RemoveAllTasks()
		}
	}
}

func (w *Worker) GetTasks() []*Task {
	return w.tasks
}

func (w *Worker) Start() error {
	if w.isRunning {
		return errors.New("worker is already running")
	}

	go w.performTasks()

	w.isRunning = true
	return nil
}

func (w *Worker) Stop() error {
	if !w.isRunning {
		return errors.New("worker is not running")
	}
	w.stopChan <- true
	w.isRunning = false

	return nil
}
