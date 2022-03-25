package main

import "fmt"

func canPartitionKSubsets(nums []int, k int) bool {
	// 排除一些基本情况
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	used := make([]bool, len(nums))
	// k号桶初始什么都没装，从nums[0]开始做选择
	// k号桶正在思考是否将nums[start]装进来，目前k号桶元素和为bucket
	// used标志某一个元素是否已经装到桶中，target是目标和
	return backtrack(k, 0, 0, target, nums, used)
}

func backtrack(k, bucket, start, target int, nums []int, used []bool) bool {
	// 所有桶都被装满了，且nums一定全部用完了
	// 因为target == sum / k
	if k == 0 {
		return true
	}
	// 装满了当前桶，递归穷举下一个桶的选择
	// 让下一个桶从nums[0]开始选数字
	if bucket == target {
		return backtrack(k-1, 0, 0, target, nums, used)
	}
	//从start开始向后探查有效的nums[i]装入当前桶
	for i := start; i < len(nums); i++ {
		// 剪枝
		if bucket+nums[i] > target { // 当前桶装不下nums[i]
			continue
		}
		if used[i] { // nums[i]已被装入别的桶中
			continue
		}
		// 做选择，将nums[i]装入桶中
		bucket += nums[i]
		used[i] = true
		// 递归穷举下一个数字是否装入当前桶
		if backtrack(k, bucket, i+1, target, nums, used) {
			return true
		}
		// 撤销选择
		bucket -= nums[i]
		used[i] = false
	}
	// 穷举所有数字，都无法装满当前桶
	return false
}

func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	k := 4
	fmt.Println(canPartitionKSubsets(nums, k))
}
