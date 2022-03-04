package main

import "fmt"

func quickSort(array []int, len int) {
	quickSortC(array, 0, len-1)
}

func quickSortC(array []int, left int, right int) {
	if left >= right {
		return
	}
	mid := partition(array, left, right)
	quickSortC(array, left, mid-1)
	quickSortC(array, mid+1, right)
}

func partition(array []int, left int, right int) int {
	pivot := array[right]
	i := left
	for j := left; j < right; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i += 1
		}
	}
	array[i], array[right] = array[right], array[i]
	return i
}

func main() {
	nums := []int{6, 11, 3, 9, 8}
	quickSort(nums, len(nums))
	fmt.Println(nums)
}
