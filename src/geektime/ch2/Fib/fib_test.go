package Fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var (
		a int = 1
		b int = 1
	)
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		temp := a
		a = b
		b = temp + a
	}
}
