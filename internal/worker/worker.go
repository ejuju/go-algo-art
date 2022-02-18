package worker

import (
	"image"
	"time"
)

// Task represents something to execute (to apply transformation to image(s) for example)
type Task struct {
	ID              string
	Assets          []image.Image
	Transformations Transformation
}

// Result is a struct that gets returned when a task is run
type Result struct {
	Err    error
	Assets []image.Image
}

// TaskFunc represents the function that gets executed when the task is run
type TaskFunc func(id string, assets []string) error

// NewTask initiates a Task struct
func NewTask(id string, assets []string, f TaskFunc) *Task {
	return &Task{
		ID:     id,
		Assets: assets,
		Func:   f,
	}
}

// Do is a function that shoudld be executed in a seperate routine
// It receives an error channel to send errors to
func (t *Task) Do(errCh chan error) {
	errCh <- t.Func(t.ID, t.Assets)
}

// Worker represents a set of tasks to execute
// A Worker has an error channel where tasks can signal errors (or return nil of no error occured)
type Worker struct {
	ID    string
	Tasks []*Task
	ErrCh chan error
}

// NewWorker initiates a Worker struct
func NewWorker(id string, tasks ...*Task) *Worker {
	errCh := make(chan error, len(tasks))

	return &Worker{
		ID:    id,
		ErrCh: errCh,
		Tasks: tasks,
	}
}

// Work runs all tasks their own go routine
// It returns possible errors and the duration taken to complete all tasks
// Work can be called in a go routine when using several workers
func (w *Worker) Work() ([]error, time.Duration) {
	start := time.Now()

	for _, t := range w.Tasks {
		go t.Do(w.ErrCh)
	}

	var errors []error

	for i := 0; i < len(w.Tasks); i++ {
		err := <-w.ErrCh
		if err != nil {
			errors = append(errors, err)
		}
	}

	dur := time.Now().Sub(start)

	return errors, dur
}
