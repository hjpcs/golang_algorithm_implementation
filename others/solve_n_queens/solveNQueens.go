package main

import "fmt"

var res [][]string

func solveNQueens(n int) [][]string {
	// 初始化结果和空棋盘
	res = [][]string{}
	board := initBoard(n)
	backtrack(board, 0)
	return res
}

func backtrack(board [][]string, row int) {
	// 触发结束条件
	if row == len(board) {
		temp := make([]string, 0)
		// 遍历棋盘每一行，作为一个string添加进temp中
		for i := 0; i < len(board); i++ {
			var s string
			for _, v := range board[i] {
				s += v
			}
			temp = append(temp, s)
		}
		// 再将整个棋盘添加进res中
		res = append(res, temp)
		return
	}
	n := len(board[row])
	for col := 0; col < n; col++ {
		// 排除不合法选择
		if !isValid(board, row, col) {
			continue
		}
		// 做选择
		board[row][col] = "Q"
		// 进行下一行决策
		backtrack(board, row+1)
		// 撤销选择
		board[row][col] = "."
	}
}
func isValid(board [][]string, row, col int) bool {
	n := len(board)
	// 检查列是否有冲突
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	// 检查右上方是否有冲突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	// 检查左上方是否有冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
func initBoard(n int) [][]string {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	return board
}

func main() {
	fmt.Println(len(solveNQueens(8)))
	fmt.Println(solveNQueens(8))
}
