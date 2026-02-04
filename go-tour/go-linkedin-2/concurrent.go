package main

//sync.Waitgroup

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go mySqrt(&wg)
	wg.Wait()
	my4thOrder()
}
func mySqrt(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(math.Sqrt(2))
}

func my4thOrder() {
	fmt.Println(math.Sqrt(math.Sqrt(2)))
}
