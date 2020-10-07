package main

import "fmt"

func main() {
	// map panic
	// var m = make(map[string]int)
	var m map[string]int
	//m1 := new(map[string]int)
	m = make(map[string]int)

	key := "two"
	m[key] = 2

	fmt.Println(m[key])
}
