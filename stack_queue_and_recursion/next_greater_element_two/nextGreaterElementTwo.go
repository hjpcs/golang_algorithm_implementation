package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	length := len(nums)
	stack := make([]int, 0)
	ans := make([]int, length*2)
	for i := length*2 - 1; i >= 0; i-- {
		num := nums[i%length]
		for len(stack) != 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = -1
		} else {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, num)
	}
	return ans[0:length]
}

func main() {
	nums := []int{1, 2, 3, 4, 3}
	fmt.Println(nextGreaterElements(nums))
}
