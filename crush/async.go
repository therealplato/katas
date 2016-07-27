package main

// async.go runs portions of the task asynchronously

// Job describes a task for the worker
// This job is to generate a Diff containing []int of width b-a containing values k
type Job struct {
	t Transform
}

// Diff is a fragment of a State slice affected by this transform
type Diff struct {
	t Transform
	s State // not full width
	l int
}

// Worker is spawned by a goroutine to help main
type Worker struct {
	input  <-chan Job
	output chan<- Diff
}

// Work loops through input
func (w *Worker) Work() {
	var in Job
	for {
		// wait for job:
		in = <-w.input
		// calculate how long the diff is:
		l := in.t.b - in.t.a + 1 // a,b one indexed and inclusive
		// make a too-big scratch pad:
		scratch := make([]int, 2*l)
		i := 0 // zero index left of populated
		j := 1 // zero index after populated k's
		z := 1 // width of this copy
		// populate scratch with k's longer than diff:
		scratch[0] = in.t.k
		for j <= l {
			copy(scratch[i+z:j+z], scratch[i:j])
			i += z
			j += z
			z += z
		}
		// return diff:
		w.output <- Diff{
			s: scratch[:l],
			t: in.t,
			l: l,
		}
	}
}
