package main

func climbStairsRecursion(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return climbStairsRecursion(n-1) + climbStairsRecursion(n-2)
	}
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	var result int
	prev1 := 1
	prev2 := 2
	for i := 3; i <= n; i++ {
		result = prev1 + prev2
		prev1 = prev2
		prev2 = result
	}
	return result
}
