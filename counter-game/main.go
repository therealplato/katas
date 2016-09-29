package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
)

func main() {
	tcs := parse()
	for _, tc := range tcs {
		play(&tc)
	}
	for _, tc := range tcs {
		fmt.Println(tc.Winner)
	}
}

type testcase struct {
	T      int
	N      uint64
	NN     []byte
	Turn   bool // false: Louise's turn
	Winner string
}

func parse() []testcase {
	tcs := make([]testcase, 0, 10)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	for scanner.Scan() {
		tc := testcase{}
		n, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		tc.T = int(t)
		tc.N = n
		tc.NN = make([]byte, 8)
		binary.BigEndian.PutUint64(tc.NN, tc.N)
		tcs = append(tcs, tc)
	}
	return tcs
}

func play(tc *testcase) {
	if tc.N == 1 {
		return
	}
	if !is2(tc) {
		sub(&tc)
	} else {
		div(&tc)
	}
	flip(&tc)
	play(&tc)
}

// func is2(x []byte) bool {
func is2(tc *testcase) bool {
	// look for exactly one Hi bit:
	found := 0
	for _, b := range tc.NN {
		if b == 1 {
			found += 1
		}
		if b > 1 {
			return false
		}
	}
	if found == 0 || found > 1 {
		return false
	}
	return true
}
