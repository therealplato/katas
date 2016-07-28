package main

// async.go runs portions of the task asynchronously

// Job describes a task for the worker
// This job is to identify if an input index is within this transform's a..b
// if  yes, return k; if no, return 0
type Job struct {
	i int // zero index to test
}

// Result may be returned over output channel after a Job is processed
type Result struct {
	i int
	k int
}

// Worker is spawned by a goroutine to help main
type Worker struct {
	t      Transform
	input  <-chan Job
	output chan<- Result
}

// Work blocks until receiving a Job on w.input channel, then returns a
// Result iff this transform affects that Job.i
func (w *Worker) Work() {
	var in Job
	for {
		// wait for job:
		in = <-w.input
		if in.i+1 >= w.t.a && in.i < w.t.b {
			w.output <- Result{
				i: in.i,
				k: w.t.k,
			}
		}
	}
}
