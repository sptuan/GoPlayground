package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 2, 5: 10}
	t.Log(m1[2])
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{1: 1}
	t.Log(m1[1])
	if v, ext := m1[2]; ext {
		t.Log("exist!")
		t.Log(v)
	} else {
		t.Log("not exist!")
		t.Log(v)
	}
	if v, ext := m1[2]; ext {
		t.Log("exist!")
		t.Log(v)
	} else {
		t.Log("not exist!")
		t.Log(v)
	}
}
