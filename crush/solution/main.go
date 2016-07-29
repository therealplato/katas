package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var N int // size of list (up to 10^7)
var M int // number of transforms (up to 2*10^5)

// Transform describes how to change the state for this line
type Transform struct {
	a int // starting index (1 indexed inclusive)
	b int // ending index (1 indexed inclusive)
	k int // amount to add to a..b this operation (up to 10^9)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); ok == true {
		N, M = handleFirstLine(scanner.Text())
	}
	state := make([]int, N+1)
	for scanner.Scan() {
		transform := handleLine(scanner.Text())
		state = updateState(state, transform)
	}
	fmt.Println(output(state))
}

func handleFirstLine(s string) (N int, M int) {
	first := strings.Split(s, " ")
	if len(first) != 2 {
		panic("wrong len of first line")
	}
	N, _ = strconv.Atoi(first[0])
	M, _ = strconv.Atoi(first[1])
	return N, M
}

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

// On each transform, we will add flags to state marking the range that the transform affects.
// Given N=5 and t={a:1,b:3,k:10}, the flagged state will look like [10, 0, 0, -10, 0]
// After adding all flags to state, the final step is to walk left to right,
// setting each element to the rolling sum of the elements to its left: [10, 10, 10, 0, 0]
func updateState(state []int, t Transform) []int {
	i, j := t.a-1, t.b-1 // convert indices to 0 index from 1 index
	state[i] += t.k
	state[j+1] -= t.k
	return state
}

func output(state []int) string {
	var sum int
	var max int
	for i, x := range state {
		sum += x
		state[i] = sum
		if sum > max {
			max = sum
		}
	}
	return strconv.Itoa(max)
}
