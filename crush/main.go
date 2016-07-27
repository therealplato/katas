package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/profile"
)

// // N is size of list (up to 10^7)
// var N int
//
// // M is number of ops (up to 2*10^5)
// var M int

// a is first index (1 indexed)
// b is second index (1 indexed)
// k is amount to add this operation (up to 10^9)
// var a, b, k int

// Cfg stores max width and max transforms
type Cfg struct {
	N int
	M int
}

var complete int

// Transform describes how to change the state for this line
type Transform struct {
	a int
	b int
	k int
}

// State is the full row, up to len 10e9
type State []int

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	var state State
	var cfg Cfg
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		// create zeroed state of width M:
		state, cfg = handleFirstLine(scanner.Text())
	}

	queue := make(chan Job)
	diffs := make(chan Diff)
	w := &Worker{
		input:  queue,
		output: diffs,
	}
	go w.Work()
	go setup(scanner, queue)
	output(state, diffs, cfg)
}

func setup(scanner *bufio.Scanner, queue chan Job) {
	for scanner.Scan() {
		transform := handleLine(scanner.Text())
		queue <- Job{
			t: transform,
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
func output(state State, diffs chan Diff, cfg Cfg) {
	for {
		d := <-diffs
		// diff := calcTransform(state, transform)
		state = applyTransform(state, d)
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
	// var i int
	// for i = 0; i < N; i++ {
	// 	state = append(state, 0)
	// }
	// return state
}

// handleLine parses one input transformation line and updates state0
func handleLine(s string) Transform {
	nums := strings.Split(s, " ")
	// a and b are one-indexed indices
	a, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])
	// k is how much this transform adds to elements of state with indices a..b inclusive
	k, _ := strconv.Atoi(nums[2])
	return Transform{
		a: a,
		b: b,
		k: k,
	}
}

func applyTransform(state State, d Diff) State {
	i := d.t.a - 1 // state index
	var j int      // diff index
	for j < d.l {
		state[i] += d.s[j]
		i++
		j++
	}
	return state
}

//
// // updateState adds k to state, indices i through j inclusive
// func updateState(state []int, i, j, k int) []int {
// 	// for x, v := range state {
// 	// 	if int(x) >= i && int(x) <= j {
// 	// 		state[int(x)] = v + k
// 	// 	}
// 	// }
//
// 	// i is zero indexed and flags the leftmost state element to bump
// 	for y, z := range state[i : j+1] {
// 		// y is zero index within the slice of state
// 		// add i to get index within state
// 		state[y+i] = z + k
// 	}
// 	return state
// }

func max(state []int) int {
	var tmp int
	for _, x := range state {
		if x > tmp {
			tmp = x
		}
	}
	return tmp
}
