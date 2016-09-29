package main

import (
	"fmt"

	"github.com/therealplato/katas/waterbuckets/chart"
)

func main() {
	chart := chart.NewChart(4, 1, 2, 4)
	if len(chart) < 3 {
		panic("nowai")
	}
	depth := chart.Fill()
	fmt.Println(depth)
}
