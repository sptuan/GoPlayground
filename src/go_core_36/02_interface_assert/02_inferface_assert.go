package main

import "fmt"

func main() {
	var s []string
	s = append(s, "string1")
	s = append(s, "string2")
	value, ok := interface{}(s).([]string)
	if ok != true {
		fmt.Println("value is not a []string")
	} else {
		fmt.Println("value is []string")
		fmt.Println(value)
	}
}
