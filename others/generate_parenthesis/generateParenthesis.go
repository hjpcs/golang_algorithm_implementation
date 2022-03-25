package main

import "fmt"

var res []string

func generateParenthesis(n int) []string {
	res = []string{}
	brackets := []string{"(", ")"}
	backtrack(2*n, brackets, "")
	return res
}

func backtrack(length int, brackets []string, s string) {
	if length == 0 {
		if isValid(s) {
			res = append(res, s)
		}
		return
	}
	for i := 0; i < len(brackets); i++ {
		s += brackets[i]
		backtrack(length-1, brackets, s)
		s = s[:len(s)-1]
	}
}

func isValid(s string) bool {
	balance := 0
	for _, char := range s {
		if char == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}

func main() {
	fmt.Println(generateParenthesis(3))
}
