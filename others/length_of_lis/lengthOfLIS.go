package length_of_lis

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	length := dp[0]
	for _, v := range dp {
		length = max(length, v)
	}
	return length
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
