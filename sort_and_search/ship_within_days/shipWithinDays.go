package main

import (
	"fmt"
)

func shipWithinDays(weights []int, days int) int {
	right := 1
	left := weights[0]
	for _, v := range weights {
		if v > left {
			left = v
		}
		right += v
	}
	for left < right {
		mid := left + (right-left)/2
		if shipWithCap(weights, mid) <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func shipWithCap(weights []int, x int) int {
	var days int
	for i := 0; i < len(weights); {
		c := x
		for i < len(weights) {
			if weights[i] > c {
				break
			} else {
				c -= weights[i]
				i++
			}
		}
		days++
	}
	return days
}

func main() {
	weights := []int{3, 2, 2, 4, 1, 4}
	days := 3
	fmt.Println(shipWithinDays(weights, days))
}
