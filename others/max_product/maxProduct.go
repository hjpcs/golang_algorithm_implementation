package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	getmin := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	getmax := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	max := math.MinInt
	imax, imin := 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = getmax(imax*nums[i], nums[i])
		imin = getmin(imin*nums[i], nums[i])
		max = getmax(max, imax)
	}
	return max
}

func main() {
	nums := []int{2, 3, -2, 4, -2}
	fmt.Println(maxProduct(nums))
}
