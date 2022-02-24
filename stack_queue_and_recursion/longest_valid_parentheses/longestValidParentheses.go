package main

func longestValidParentheses(s string) int {
	left, right, length := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left < right {
			left, right = 0, 0
		}
		if left == right {
			length = max(left+right, length)
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if left > right {
			left, right = 0, 0
		}
		if left == right {
			length = max(left+right, length)
		}
	}
	return length
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
