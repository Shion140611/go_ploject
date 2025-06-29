package main

import "fmt"

func main() {

	fmt.Println(sumSlice([]int{1, 2, 3}))
}

func sumSlice(nam []int) int {
	sum := 0
	for _, number := range nam {
		sum += number
	}
	return sum
}
