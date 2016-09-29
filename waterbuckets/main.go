package main

import "fmt"

type Chart []element

type element struct {
	chart  *Chart
	index  int
	height int
}

type window struct {
	chart  Chart
	width  int
	leftI  int // right index of window, inclusive
	rightI int
}

func main() {
	chart := newChart(4, 1, 2, 4)
	fmt.Println(chart)
	// depth := fill(chart)
}

func fill(c Chart) (depth int) {
	return depth
}

func newChart(heights ...int) Chart {
	c := Chart{}
	for i, h := range heights {
		fmt.Printf("i: %d, h: %d\n", i, h)
		e := element{
			chart:  &c,
			index:  i,
			height: h,
		}
		c = append(c, e)
	}
	return c
}
