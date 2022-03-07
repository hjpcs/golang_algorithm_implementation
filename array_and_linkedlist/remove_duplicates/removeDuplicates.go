package remove_duplicates

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[slow] == nums[fast] {
			fast++
		} else {
			slow++
			nums[slow] = nums[fast]
			fast++
		}
	}
	return slow + 1
}
