package main

import (
	"fmt"
)

const A = byte('a')
const B = byte('b')
const C = byte('c')
const D = byte('d')

func main() {
	c := bake([]byte("abc"))
	// q.Q(c)
	fmt.Println(c)
	// sift(c) will find palindromes available from this root
	// Sweep the root across and sift each
	// Accumulate palindroms rooted at each character
}

// Start at the top layer.
// Add one if palindromic.
// Build a tree testing (this layer + left leaf + right leaf)  for palindromes.
// If the pair of left and right leaf nodes are not the same, delete nodes until a palindrome is found.
// Test all paths.
func sift(c *cake) (n int) {
	d := c.copy()
	i := *d.top()
	if pal(i.val()) {
		n++
	}
	for _, x := range []byte{A, B, C, D} {
		for i.chars[i.i] != x {
			d.top().down = d.top().down.down
			// delete left
		}
		for i.chars[i.j] != x {
			// delete right
		}
	}
	return n
}
