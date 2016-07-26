package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/profile"
)

// N is size of list (up to 10^7)
var N int

// M is number of ops (up to 2*10^5)
var M int

// a is first index (1 indexed)
// b is second index (1 indexed)
// k is amount to add this operation (up to 10^9)
var a, b, k int

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	var state []int
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		state = handleFirstLine(scanner.Text())
	}
	for scanner.Scan() {
		state = handleLine(state, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(max(state))
}

func handleFirstLine(s string) []int {
	first := strings.Split(s, " ")
	if len(first) != 2 {
		panic("wrong len of first line")
	}
	N, _ = strconv.Atoi(first[0])
	M, _ = strconv.Atoi(first[1])
	// initialize zeroed array:
	return make([]int, N)
	// var i int
	// for i = 0; i < N; i++ {
	// 	state = append(state, 0)
	// }
	// return state
}

// handleLine parses one input transformation line and updates state0
func handleLine(state []int, s string) []int {
	nums := strings.Split(s, " ")
	// a and b are one-indexed indices
	a, _ = strconv.Atoi(nums[0])
	b, _ = strconv.Atoi(nums[1])
	// k is how much this transform adds to elements of state with indices a..b inclusive
	k, _ = strconv.Atoi(nums[2])
	return updateState(state, a-1, b-1, k)
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
