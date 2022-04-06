package min_path_sum

func minPathSum(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	for i := 1; i < column; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < row; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[row-1][column-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
