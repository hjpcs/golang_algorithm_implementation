package main

import (
	"fmt"
)

func firstMissingPositiveByMap(nums []int) int {
	n := len(nums)
	valueMap := make(map[int]int)
	for i := 0; i < n; i++ {
		valueMap[nums[i]]++
	}
	for i := 1; i <= n; i++ {
		if _, ok := valueMap[i]; !ok {
			return i
		}
	}
	return n + 1
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func main() {
	num1 := []int{1, 2, 0}
	num2 := []int{3, 4, -1, 1}
	num3 := []int{7, 8, 9, 11, 12}
	fmt.Println(firstMissingPositive(num1))
	fmt.Println(firstMissingPositive(num2))
	fmt.Println(firstMissingPositive(num3))
}
