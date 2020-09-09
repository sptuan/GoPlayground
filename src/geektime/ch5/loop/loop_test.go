package loop_test

import "testing"

func TestLoop(t *testing.T) {
	a := 5
	for a < 10 {
		a++
		t.Log(a)
	}
}
