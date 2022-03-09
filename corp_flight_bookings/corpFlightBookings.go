package main

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)
	for _, booking := range bookings {
		diff[booking[0]-1] += booking[2]
		if booking[1] != n {
			diff[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	return diff
}

func main() {
	bookings := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 5, 25},
	}
	n := 5
	fmt.Println(corpFlightBookings(bookings, n))
}
