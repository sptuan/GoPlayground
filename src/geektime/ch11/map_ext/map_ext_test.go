package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[3](2))
}

func TestMapForSet(t *testing.T) {
	s := map[int]bool{}
	s[1] = true
	t.Log(len(s))
	s[1] = false
	t.Log(len(s))
	delete(s, 1)
	t.Log(len(s))

}
