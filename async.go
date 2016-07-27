package main

// async.go runs portions of the task asynchronously

var Zeroed State

type Job struct {
	State     *State
	Transform Transform
}

// Worker is spawned by a goroutine to help main
type Worker struct {
	input   <-chan Job
	scratch State
	output  chan<- State
}

// Work loops through input
func (w *Worker) Work() {
	for {
		in := <-w.input
		copy(w.scratch, *in.State)
		applyTransform(w.scratch, in.Transform)
		output <- w.scratch
	}
}
