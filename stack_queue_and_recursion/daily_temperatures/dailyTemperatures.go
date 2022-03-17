package main

import "fmt"

func dailyTemperaturesFront(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	var stack []int
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	var stack []int
	for i := length - 1; i >= 0; i-- {
		t := temperatures[i]
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] <= t {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = 0
		} else {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	t := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(t))
}
