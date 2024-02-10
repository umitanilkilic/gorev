package gorev

import (
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	TaskInterface TaskInterface
	Priority      int
	TaskID        uint32
	//TODO: timeout sistemi ekle
	//TimeOut       time.Duration
}

type TaskInterface interface {
	Perform() error
}

func NewTask(task TaskInterface, priority int) (*Task, error) {
	if task == nil {
		return nil, fmt.Errorf("task is nil")
	}
	//return &Task{TaskInterface: task, Priority: priority, TimeOut: timeOut, TaskID: uuid.New().ID()}, nil
	return &Task{TaskInterface: task, Priority: priority, TaskID: uuid.New().ID()}, nil
}