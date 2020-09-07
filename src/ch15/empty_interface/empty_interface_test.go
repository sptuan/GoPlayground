package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//if i,ok := p.(int); ok{
	//	fmt.Println("Interger", i)
	//	return
	//}
	//if s,ok:= p.(string); ok{
	//	fmt.Println("String", s)
	//	return
	//}
	//fmt.Println("Unknown Type")
	switch p.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown Type")
	}
}

func TestInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("hahahaha")
}
