package main

import (
	"flag"
	"fmt"
)

var name string
var age int
var gender string

func init() {
	flag.StringVar(&name, "name", "NULL", "Name of the person.")
	flag.IntVar(&age, "age", -1, "Age of the person.")
	flag.StringVar(&gender, "gender", "unknown", "Gender of the person.")
}

func main() {
	flag.Parse()

	fmt.Println(name, age, gender)
}
