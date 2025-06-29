package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println("isEven(1):", isEven(1)) // false
	fmt.Println("isEven(2):", isEven(2)) // true
}
