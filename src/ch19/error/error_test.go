package err

import (
	"errors"
	"testing"
)

//区分错误类型
var LessThanTwoError = errors.New("N less than 0")
var LargerThanTwoError = errors.New("N larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 0 || n > 100 {
		return nil, LessThanTwoError
	}
	return nil, nil
}

func TestGet(t *testing.T) {
	if _, err := GetFibonacci(-10); err != nil {
		t.Error(err)
		if err == LessThanTwoError {
			t.Log("LessThanError")
		}
	} else {
		t.Log(100)
	}
}
