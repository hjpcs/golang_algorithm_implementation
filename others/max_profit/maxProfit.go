package main

import "fmt"

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		profit := dp[i-1] + prices[i] - prices[i-1]
		if profit > 0 {
			dp[i] = profit
		}
	}
	max := dp[0]
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
