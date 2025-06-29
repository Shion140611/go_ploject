package main

import (
	"fmt"
)

func main() {
	printNumbers()
}

func printNumbers() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}
