package imgutil

import (
	"image"
	"time"
)

// deprecated
type TransformationFunc func(image.Image) image.Image

// deprecated
func Transform(img image.Image, funcs ...TransformationFunc) image.Image {
	for _, f := range funcs {
		img = f(img)
	}
	return img
}

// Worker runs multiple tasks concurrenlty
type Worker struct {
	ID        string
	Tasks     []Task
	ResultsCh chan Result
}

//
type Task struct {
	ID       string
	Input    image.Image
	Pipeline []TaskFunc
}

// TaskFunc represents a function that does something with an image (for example: resize it)
// It returns a new image and a possible error
type TaskFunc func(input image.Image) (image.Image, error)

//
func (t *Task) Do(outCh chan Result) {
	var err error
	var res Result
	var img = t.Input

	for _, f := range t.Pipeline {
		img, err = f(img)
		if err != nil {
			break
		}
	}

	res.Err = err
	res.Out = img

	outCh <- res
}

// Result is a struct that gets returned when a an edit is run is run
type Result struct {
	Err error
	Out image.Image
}

// Work runs all tasks their own go routine
// It returns possible errors and the duration taken to complete all tasks
// Work can be called in a go routine when using several workers
func (w *Worker) Work() ([]Result, time.Duration) {
	var start = time.Now()

	for _, task := range w.Tasks {
		go task.Do(w.ResultsCh)
	}

	var results []Result

	for range w.Tasks {
		res := <-w.ResultsCh
		results = append(results, res)
	}

	var dur = time.Now().Sub(start)

	return results, dur
}

// Do is a function that shoudld be executed in a seperate routine
// It receives an error channel to send errors to
// func (t *Task) Do(errCh chan error) {
// 	errCh <- t.Func(t.ID, t.Assets)
// }

// Worker represents a set of tasks to execute
// A Worker has an error channel where tasks can signal errors (or return nil of no error occured)
// type Worker struct {
// 	ID    string
// 	Tasks []*Task
// 	ErrCh chan error
// }

// NewWorker initiates a Worker struct
// func NewWorker(id string, tasks ...*Task) *Worker {
// 	errCh := make(chan error, len(tasks))

// 	return &Worker{
// 		ID:    id,
// 		ErrCh: errCh,
// 		Tasks: tasks,
// 	}
// }
