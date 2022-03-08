package main

import "fmt"

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}
	return NumArray{preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right+1] - this.sum[left]
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	preSum := Constructor(nums)
	fmt.Println(preSum.SumRange(0, 2))
	fmt.Println(preSum.SumRange(2, 5))
	fmt.Println(preSum.SumRange(0, 5))
}
