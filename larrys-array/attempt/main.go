package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// T is number of testcases
var T int

// Testcase is one input to test
type Testcase struct {
	N int   // size of A
	A []int // starting index (1 indexed inclusive)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); ok == true {
		T, _ = strconv.Atoi(scanner.Text())
	}
	tcs := make(chan Testcase)
	go read(scanner, tcs)
	for i := 0; i < T; i++ {
		tc := <-tcs
		if solvable(tc) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func read(scanner *bufio.Scanner, tcs chan Testcase) {
	flagOdd := true
	var tmpN int
	for scanner.Scan() {
		if flagOdd {
			tmpN, _ = strconv.Atoi(scanner.Text())
		} else {
			tc := Testcase{
				N: tmpN,
				A: make([]int, tmpN),
			}
			aa := strings.Split(scanner.Text(), " ")
			for i, a := range aa {
				tc.A[i], _ = strconv.Atoi(a)
			}
			tcs <- tc
		}
		flagOdd = !flagOdd
	}
}

func solvable(tc Testcase) bool {
	return false
}
