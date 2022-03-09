package car_pooling

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)
	for _, trip := range trips {
		diff[trip[1]] += trip[0]
		diff[trip[2]] -= trip[0]
	}
	preSum := 0
	for _, v := range diff {
		preSum += v
		if preSum > capacity {
			return false
		}
	}
	return true
}
