package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	res := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[nums2[i]] = -1
		} else {
			res[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for i := 0; i < len(nums1); i++ {
		nums1[i] = res[nums1[i]]
	}
	return nums1
}

func main() {
	num1 := []int{4, 1, 2}
	num2 := []int{1, 3, 4, 2}
	res := nextGreaterElement(num1, num2)
	fmt.Println(res)
}
