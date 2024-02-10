package gorev

import (
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	TaskInterface Tasker
	Priority      int
	taskID        uint32
	//TODO: timeout sistemi ekle
	//TimeOut       time.Duration
}

type Tasker interface {
	Perform() error
}

func NewTask(task Tasker, priority int) (*Task, error) {
	if task == nil {
		return nil, fmt.Errorf("task is nil")
	}
	//return &Task{TaskInterface: task, Priority: priority, TimeOut: timeOut, TaskID: uuid.New().ID()}, nil
	return &Task{TaskInterface: task, Priority: priority, taskID: uuid.New().ID()}, nil
}

func (t *Task) GetTaskID() uint32 {
	return t.taskID
}
