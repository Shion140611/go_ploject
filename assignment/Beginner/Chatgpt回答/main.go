package main

import (
	"fmt"
)

// 偶数判定関数
func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println(isEven(4)) // true
	fmt.Println(isEven(5)) // false
}
