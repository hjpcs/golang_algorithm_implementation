package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	getmin := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	row := len(triangle)
	for i := 1; i < row; i++ {
		triangle[i][0] += triangle[i-1][0]
		triangle[i][i] += triangle[i-1][i-1]
	}
	for i := 2; i < row; i++ {
		for j := 1; j < len(triangle[i])-1; j++ {
			triangle[i][j] += getmin(triangle[i-1][j-1], triangle[i-1][j])
		}
	}
	min := triangle[row-1][0]
	for _, v := range triangle[row-1] {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
