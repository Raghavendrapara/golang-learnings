package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "abv"
	a[1] = "abcc"
	fmt.Println(len(a))

	p := [2]string{"", ""}
	fmt.Println(p)

	//Note slices are references to array not new arrays per se
	names := [5]string{"ina", "mina", "dika", "amar", "akbar"}

	fmt.Println(names)
	c := names[0:2]
	d := names[2:5]
	fmt.Println(c, d)
	c[1] = "jane"
	fmt.Println(c)
	fmt.Println(names)

}
