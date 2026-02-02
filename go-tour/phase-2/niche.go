//defer

package main

import (
	"fmt"
)

type cood struct {
	X int
	Y int
}

func main() {
	play()
	playPointers()
	v := cood{1, 5}
	fmt.Printf("%T ", v)
	v.X = 5
	fmt.Println(cood{1, 2})

	// Can access the pointer without using *
	p := &v
	// fmt.Println((*p).X)
	p.X = 6
	fmt.Println(p.X)
	// defer fmt.Println("world from main")
	// fmt.Println("hello from main")
}
func play() {
	// defer fmt.Println("world")
	// fmt.Println("hello")
}

func playPointers() {
	i, j := 1, 2
	fmt.Println(i, j)

	p := &i
	fmt.Println(i, "   ", p, "  ", *p)

	*p = 33
	fmt.Println(i)

}
