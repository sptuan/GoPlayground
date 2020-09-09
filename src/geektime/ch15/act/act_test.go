package act

import (
	"fmt"
	"testing"
)

func changeVal(n *int) {
	*n = 10
	fmt.Println(*n)
}

func TestAct(t *testing.T) {
	test := 20
	t.Log(test)
	changeVal(&test)
	t.Log(test)
}

type Employee struct {
	id   int
	name string
	age  int
}

func (e *Employee) String() string {
	return fmt.Sprintf("ID:%d/Name:%s/Age:%d", e.id, e.name, e.age)
}

//func (e Employee) String() string {
//	return fmt.Sprintf("ID:%d/Name:%s/Age:%d", e.id, e.name, e.age)
//}

func TestStructOperations(t *testing.T) {
	e := &Employee{id: 10, name: "BOB", age: 20}
	t.Log(e.String())
}
