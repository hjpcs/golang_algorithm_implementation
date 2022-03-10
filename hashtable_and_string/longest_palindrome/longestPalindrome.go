package main

import "fmt"

func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		s1 := Palindrome(s, i, i)
		s2 := Palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func Palindrome(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

func main() {
	s := "xaabbaay"
	fmt.Println(longestPalindrome(s))
}
