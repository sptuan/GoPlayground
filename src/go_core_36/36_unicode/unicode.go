package main

import "fmt"

func main() {
	str := "Go语言，这是一个测试用字符串"
	fmt.Printf("string: %q\n", str)
	fmt.Printf("-> runes(char): %q\n", []rune(str))
	fmt.Printf("-> runes(hex): %x\n", []rune(str))
	fmt.Printf("-> bytes(hex): [% x]\n", []byte(str))

	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}
}
