package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args {
		fmt.Println(idx, arg)
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			fmt.Print(input.Text())
		}
	}
}
