package main

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1000000001
	for left < right {
		mid := left + (right-left)/2
		if eatingTime(piles, mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func eatingTime(piles []int, speed int) int {
	var h int
	for i := 0; i < len(piles); i++ {
		h += piles[i] / speed
		if piles[i]%speed != 0 {
			h++
		}
	}
	return h
}
