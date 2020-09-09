package array_test

import "testing"

func TestArray(t *testing.T) {
	var arr [3]int
	arr1 := [...]int{1, 2, 3, 4}
	t.Log(arr[0], arr[1], arr[2])
	t.Log(arr1)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{2, 3, 4, 5, 6, 7}
	arr4 := arr3[2:]
	arr5 := arr3[2:4]
	t.Log(arr3)
	t.Log(arr4)
	t.Log(arr5)
}
