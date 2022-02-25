package main

import (
	"fmt"
)

func mySqrt(x int) int {
	left, right := 0, x
	ans := -1
	for left <= right {
		mid := (left + right) >> 1
		if mid*mid > x {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}
	return ans
}

func main() {
	x := 36
	fmt.Println(mySqrt(x))
}
