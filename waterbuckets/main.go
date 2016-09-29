package main

import "fmt"

type Chart []int

type window struct {
	chart  Chart
	width  int
	leftI  int // right index of window, inclusive
	rightI int
}

func main() {
	chart := newChart(4, 1, 2, 4)
	if len(chart) < 3 {
		panic("nowai")
	}
	depth := fill(chart)
	fmt.Println(depth)
}

func newChart(heights ...int) Chart {
	c := Chart{}
	for _, h := range heights {
		c = append(c, h)
	}
	return c
}

// fill takes a chart and finds the total depth of all holes in the chart
// i.e. all regions bounded by a higher height somewhere to the left and a
// higher hight somewhere to the right
func fill(c Chart) (depth int) {
	for width := 3; width <= len(c); width++ {
		filled := fillWindowedHoles(&c, width)
		depth += filled
	}
	return depth
}

// fillWindowedHoles sweeps a window of width w across the chart
// if any element within the window is shorter than the left and right heights
// of the window, the element is filled in to that height (modifying *c). The
// total depth of filled holes is returned.
func fillWindowedHoles(c *Chart, w int) int {
	accumulator := 0
	// Determine the left index of the rightmost window:
	finalLeftIndex := len(*c) - w
	for i0 := 0; i0 <= finalLeftIndex; i0++ {
		// Find the right index of the current window:
		i1 := i0 + w - 1
		// Find the left and right heights of the window:
		x0 := (*c)[i0]
		x1 := (*c)[i1]
		// Find the smaller of the left and right bounds:
		minBound := x0
		if x1 < x0 {
			minBound = x1
		}
		// Check each height inside the window, looking for holes of smaller height
		// than the minimum bound:
		for i := i0 + 1; i <= i1-1; i++ {
			// Calculate the difference from the minimum bound of the window:
			d := minBound - (*c)[i]
			if d > 0 {
				// there is a hole, fill it to the min bound of this window:
				(*c)[i] += d
				accumulator += d
			}
		}
	}
	return accumulator
}
