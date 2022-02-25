package main

import "fmt"

func mySqrt(x int) int {
	left, right := 0, x
	mid := (left + right) >> 1
	for left <= right {
		if mid*mid > x {
			right = mid - 1
			mid = (left + right) >> 1
		} else if mid*mid < x {
			left = mid + 1
			mid = (left + right) >> 1
		} else {
			return mid
		}
	}
	if mid*mid < x {
		return mid
	}
	return mid + 1
}

func main() {
	x := 8
	fmt.Println(mySqrt(x))
}
