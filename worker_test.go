package gorev

import (
	"errors"
	"testing"
)

func TestNewWorker(t *testing.T) {
	w := NewWorker()
	if w.workerId == 0 {
		t.Errorf("WorkerId is 0")
	}
	if w.stopChan == nil {
		t.Errorf("stopChan is nil")
	}
}

type testTask struct{}

func (t *testTask) Perform() error {
	return errors.New("test error")
}

func TestAddTask(t *testing.T) {
	w := NewWorker()
	t1, _ := NewTask(&testTask{}, 1)
	w.AddTask(t1)
	if len(w.tasks) != 1 {
		t.Errorf("Task not added")
	}
}

func TestSortTasks(t *testing.T) {
	w := NewWorker()
	t1, _ := NewTask(&testTask{}, 1)
	t2, _ := NewTask(&testTask{}, 2)
	w.AddTask(t1)
	w.AddTask(t2)
	if w.tasks[0].Priority != 2 {
		t.Errorf("Tasks not sorted")
	}
}

func TestRemoveTaskByIndex(t *testing.T) {
	w := NewWorker()
	t1, _ := NewTask(&testTask{}, 1)
	t2, _ := NewTask(&testTask{}, 2)
	w.AddTask(t1)
	w.AddTask(t2)
	w.RemoveTaskByIndex(0)
	if len(w.tasks) != 1 {
		t.Errorf("Task not removed")
	}
}

func TestRemoveAllTask(t *testing.T) {
	w := NewWorker()
	t1, _ := NewTask(&testTask{}, 1)
	t2, _ := NewTask(&testTask{}, 2)
	w.AddTask(t1)
	w.AddTask(t2)
	w.RemoveAllTasks()
	if len(w.tasks) != 0 {
		t.Errorf("Tasks not removed")
	}
}

func TestRemoveTaskByTaskID(t *testing.T) {
	w := NewWorker()
	t1, _ := NewTask(&testTask{}, 1)
	t2, _ := NewTask(&testTask{}, 2)
	w.AddTask(t1)
	w.AddTask(t2)
	w.RemoveTaskByTaskID(t1.GetTaskID())
	if len(w.tasks) != 1 {
		t.Errorf("Task not removed")
	}
}

func TestGetTasks(t *testing.T) {
	w := NewWorker()
	t1, _ := NewTask(&testTask{}, 1)
	w.AddTask(t1)
	tasks := w.GetTasks()
	if len(tasks) != 1 {
		t.Errorf("Task not returned")
	}
}

func TestStartWorker(t *testing.T) {
	w := NewWorker()
	t1, _ := NewTask(&testTask{}, 1)
	t2, _ := NewTask(&testTask{}, 2)
	w.AddTask(t1)
	w.AddTask(t2)
	err := w.Start()
	if err != nil {
		t.Errorf("starting function not working properly")
	}
}

func TestStopWorker(t *testing.T) {
	w := NewWorker()
	err := w.Stop()
	if err == nil {
		t.Errorf("stopping function not working properly")
	}

	err = w.Start()
	if err != nil {
		t.Errorf("starting function not working properly")
	}
	err = w.Stop()
	if err != nil {
		t.Errorf("stopping function not working properly")
	}
}
