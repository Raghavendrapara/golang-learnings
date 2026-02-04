package main

import "fmt"

func main() {
	users := map[string]string{"admin": "root"}
	tryToReplace(users)

	fmt.Printf("Address of map m: %p\n", &users)
	fmt.Println(*&users)
	addressPrint(&users)
}

func tryToReplace(m map[string]string) {
	m = make(map[string]string)
	m["hack"] = "attack"
}

func addressPrint(ptr *map[string]string) {
	fmt.Println(&ptr)
	fmt.Println(*ptr)
}
