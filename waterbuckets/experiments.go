package main

// Problem:
// Take an array of integers. Consider it as describing a histogram or set of
// adjacent blocks.

// Imagine pouring water into the top of the graph; find the area filled by the
// water:

// 4 2 1 3:
//                 //
//       //    >>  // XX XX //  >> answer is "3 units contained by bounded blocks"
// //    //        // // XX //
// // // //        // // // //
// import "fmt"
//
// func main() {
// 	fmt.Println("vim-go")
// 	c := &chart{4, 2, 1, 4}
// 	if len(*c) < 3 {
// 		panic("noway")
// 	}
//
// 	fmt.Println(c.isHoleBetween(0, 3))
// }
//
// func (c *chart) checkWindow(width int) {
// 	for w := 2; w <= len(*c)-1; w++ {
// 		// Use this window width; look for holes that match this width across the graph
// 		i0, i1 := 0, w
// 		for i1 <= w {
// 			hole := c.isHoleBetween(i0, i1)
// 			if hole {
// 				fmt.Printf("hole detected between %d..%d\n", i0, i1)
// 				i0++
// 				i1++
// 			}
// 		}
// 	}
// }
//
// type chart []int
//
// func (c *chart) capacity() int {
// 	return 0
// }
//
// func (c *chart) isHoleBetween(i0, i1 int) bool {
// 	l := (*c)[i0]
// 	r := (*c)[i1]
//
// 	for i := i0 + 1; i < i1; i++ {
// 		x := (*c)[i]
// 		if x < l && x < r {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func (c *chart) findGlobalMaxima() (max int, indices []int) {
// 	// Find biggest element:
// 	maxX := 0
// 	for _, x := range *c {
// 		if x > maxX {
// 			maxX = x
// 		}
// 	}
// 	// Find all indices of elements with the biggest value:
// 	for i, x := range *c {
// 		if x == maxX {
// 			indices = append(indices, i)
// 		}
// 	}
// 	return maxX, indices
// }
//
// func (c *chart) findLocalMinima() (indices []int) {
// 	return indices
// }
