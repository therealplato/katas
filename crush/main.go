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

var complete int

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	var cfg Cfg
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		cfg = handleFirstLine(scanner.Text())
	}

	queue := make(chan Job) // input channel
	results := make(chan Result)
	done := make(chan int)
	sinks := makeSinks(cfg.M) // per-worker copies of the input channel
	go mux(queue, sinks)
	makeWorkers(scanner, sinks, results)
	go processResults(results, cfg, done)
	for i := 0; i < cfg.N; i++ {
		// ask all transforms how much they change this index
		queue <- Job{
			i: i,
		}
	}
	answer := <-done // blocks
	fmt.Println(answer)
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
		txt := scanner.Text()
		transform := handleLine(txt)
		w := &Worker{
			t:      transform,
			input:  sinks[line],
			output: results,
		}
		go w.Work()
		line++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
func processResults(results chan Result, cfg Cfg, done chan int) {
	var count int64
	var expected int64 = int64(cfg.M) * int64(cfg.N)
	var biggest int
	state := make([]int, cfg.N)
	for {
		r := <-results
		count++
		state[r.i] += r.k
		if state[r.i] > biggest {
			biggest = state[r.i]
			fmt.Printf("new max: %d\n", biggest)
		}
		if count == expected {
			done <- biggest
		}
	}
}

func handleFirstLine(s string) Cfg {
	first := strings.Split(s, " ")
	if len(first) != 2 {
		panic("wrong len of first line")
	}

	c := Cfg{}
	c.N, _ = strconv.Atoi(first[0])
	c.M, _ = strconv.Atoi(first[1])
	return c
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
