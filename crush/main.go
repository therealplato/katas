package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/profile"
)

// We will create M workers, each responsible for one transform.
// We will create initial state then iterate along the N state indices.
// For each index we will ask all workers how much they transform this index.

// Cfg stores max width and max transforms
type Cfg struct {
	N int // size of list (up to 10^7)
	M int // number of transforms (up to 2*10^5)
}

// Transform describes how to change the state for this line
type Transform struct {
	a int // first index (1 indexed)
	b int // second index (1 indexed)
	k int // amount to add this operation (up to 10^9)
}

// State represen
type State []int

var complete int

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	var state State
	var cfg Cfg
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		state, cfg = handleFirstLine(scanner.Text()) // create zeroed state of width N
	}

	queue := make(chan Job) // input channel
	results := make(chan Result)
	sinks := makeSinks(cfg.M) // per-worker copies of the input channel
	go mux(queue, sinks)
	go makeWorkers(scanner, sinks, results)

	for i := 0; i < cfg.N; i++ {
		queue <- Job{
			i: transform,
		}
	}
	processResults(state, results, cfg)
}

func makeSinks(m int) []chan Job {
	sinks := [](chan Job){}
	for i := 0; i < m; i++ {
		sinks = append(sinks, make(chan Job))
	}
	return sinks
}
func mux(source chan Job, sinks []chan Job) {
	var j Job
	for {
		j = <-source
		for _, sink := range sinks {
			sink <- j
		}
	}
}

func makeWorkers(scanner *bufio.Scanner, sinks []chan Job, results chan Result) {
	line := 0
	for scanner.Scan() {
		transform := handleLine(scanner.Text())
		w := &Worker{
			t:      transform,
			input:  sinks[line],
			output: results,
		}
		go w.Work()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
func processResults(state State, diffs chan Result, cfg Cfg) {
	for {
		_ = <-diffs
		// diff := calcTransform(state, transform)
		// state = applyTransform(state, d)
		complete++
		if complete == cfg.N {
			break
		}
	}
	fmt.Println(max(state))
}

func handleFirstLine(s string) (State, Cfg) {
	first := strings.Split(s, " ")
	if len(first) != 2 {
		panic("wrong len of first line")
	}

	c := Cfg{}
	c.N, _ = strconv.Atoi(first[0])
	c.M, _ = strconv.Atoi(first[1])
	// initialize zeroed array:
	return State(make([]int, c.N)), c
}

// handleLine parses one input transformation line and updates state0
func handleLine(s string) Transform {
	nums := strings.Split(s, " ")
	a, _ := strconv.Atoi(nums[0]) // a and b are one-indexed indices
	b, _ := strconv.Atoi(nums[1])
	k, _ := strconv.Atoi(nums[2]) // k is how much this transform adds to elements of state with indices a..b inclusive
	return Transform{
		a: a,
		b: b,
		k: k,
	}
}

func max(state []int) int {
	var tmp int
	for _, x := range state {
		if x > tmp {
			tmp = x
		}
	}
	return tmp
}
