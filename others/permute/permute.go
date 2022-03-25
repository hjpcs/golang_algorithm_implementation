package main

import "fmt"

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	track := make([]int, 0)
	backtrack(nums, track)
	return res
}

func backtrack(nums, track []int) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}
		track = append(track, nums[i])
		backtrack(nums, track)
		track = track[:len(track)-1]
	}
}

func contains(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{0, 1}
	fmt.Println(permute(nums))
}
