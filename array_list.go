package main

import (
	"fmt"
)

func fill(array [10]int) [10]int {
	for i := 1; i <= 10; i++ {
		array[i-1] = i
	}
	return array
}

func out(array [10]int) {
	fmt.Println(array)
}

func main() {
	var array [10]int
	array = fill(array)
	out(array)
}
