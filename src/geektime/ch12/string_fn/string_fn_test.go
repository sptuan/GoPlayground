package string_fn

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestingConv(t *testing.T) {
	t.Log(strconv.Itoa(234))
	s := "123552"
	if s_i, err := strconv.Atoi(s); err == nil {
		t.Log(s_i)
	}
}
