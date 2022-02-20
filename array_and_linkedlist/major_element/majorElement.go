package main

import "fmt"

func majorityElement(nums []int) int {
	valueMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		valueMap[nums[i]]++
	}
	//fmt.Println(valueMap)
	for k, v := range valueMap {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}

func boyerMoore(nums []int) int {
	count := 0
	candidate := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		}
		if nums[i] != candidate {
			count--
		}
	}
	return candidate
}

func main() {
	nums1 := []int{3, 2, 3}
	nums2 := []int{1, 7, 2, 7, 3, 7, 4, 7, 5, 7, 7}
	fmt.Println(majorityElement(nums1))
	fmt.Println(majorityElement(nums2))
	fmt.Println(boyerMoore(nums1))
	fmt.Println(boyerMoore(nums2))
}
