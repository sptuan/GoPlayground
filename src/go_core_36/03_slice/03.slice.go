package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	s3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s1)
	fmt.Println(s2)
	s1[1] = 99
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
