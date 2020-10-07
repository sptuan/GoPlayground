package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	// list
	l := list.New()
	l.PushBack("123")
	l.PushFront("000")
	for it := l.Front(); it != nil; it = it.Next() {
		fmt.Println(it.Value)
	}
	// ring
	r := ring.New(7)
	for i := 0; i < 10; i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 0; i < 15; i++ {
		fmt.Println(r.Value)
		r = r.Next()
	}
}
