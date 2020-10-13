package main

import "fmt"

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1。
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	fmt.Println("Duck type test->")
	// 示例2。
	var pet Pet = &dog
	dog.SetName("Modified 1.")
	fmt.Println(pet.Name())

	var pet2 Pet = dog
	dog.SetName("Modified 2.")
	fmt.Println(pet2.Name())
}
