package main

import "fmt"

func main() {
	c := bake([]byte("abc"))
	// var i = c.bottom()
	// for ; i != nil; i = i.up {
	// 	_ = i
	// }
	// c = shake(c)
	fmt.Printf("%s\n", c)
}
