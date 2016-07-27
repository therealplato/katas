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

type Cfg struct {
	N int
	M int
}

var cfg Cfg

type Transform struct {
	a int
	b int
	k int
}

type State []int
type TransformResult []int

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	var state State
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		state = handleFirstLine(scanner.Text())
	}
	chan queue Job
	chan diffs State
	w := &Worker{
		input: queue,

	for scanner.Scan() {
		transform := handleLine(state, scanner.Text())
		diff := calcTransform(state, transform)
		state = applyTransform(state, diff)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(max(state))
}

func handleFirstLine(s string) State {
	first := strings.Split(s, " ")
	if len(first) != 2 {
		panic("wrong len of first line")
	}

	cfg.N, _ = strconv.Atoi(first[0])
	cfg.M, _ = strconv.Atoi(first[1])
	// initialize zeroed array:
	return State(make([]int, cfg.N))
	// var i int
	// for i = 0; i < N; i++ {
	// 	state = append(state, 0)
	// }
	// return state
}

// handleLine parses one input transformation line and updates state0
func handleLine(state State, s string) Transform {
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
	// updateState(state, a-1, b-1, k)
}

func calcTransform(state State, t Transform) State {
	return updateState(state, t.a-1, t.b-1, t.k)
}

func applyTransform(state State, t State) State {
	return t
}

// updateState adds k to state, indices i through j inclusive
func updateState(state []int, i, j, k int) []int {
	// for x, v := range state {
	// 	if int(x) >= i && int(x) <= j {
	// 		state[int(x)] = v + k
	// 	}
	// }

	// i is zero indexed and flags the leftmost state element to bump
	for y, z := range state[i : j+1] {
		// y is zero index within the slice of state
		// add i to get index within state
		state[y+i] = z + k
	}
	return state
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
