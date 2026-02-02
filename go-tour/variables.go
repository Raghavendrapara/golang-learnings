package main

import (
	"fmt"
)

var d, e bool

func main() {
	a := 10
	b := 10
	fmt.Println(a, b, d, e)
	fmt.Println(add(a, b))
}
func add(a, b int) (x int) {
	return a + b
}
