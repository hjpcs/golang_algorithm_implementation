package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	s = preProcess(s)
	wordArray := strings.Split(s, " ")
	i := 0
	l := len(wordArray) - 1
	for i < l {
		wordArray[i], wordArray[l] = wordArray[l], wordArray[i]
		i++
		l--
	}
	res := strings.Join(wordArray, " ")
	return res
}

func preProcess(s string) string {
	l := len(s)
	var res []byte
	flag := 1                     // 用于处理多个连续空格
	for l != 0 && s[l-1] == ' ' { // 处理字符串后面的空格
		l--
	}
	for i := 0; i < l; i++ {
		// 添加字母后flag=0，表示允许添加空格
		if s[i] != ' ' {
			res = append(res, s[i])
			flag = 0
		}
		// flag初始化为1，可以去除字符串前面的空格
		// 添加空格后flag=1，表示不允许添加空格，即保证单词间只有一个空格
		if s[i] == ' ' && flag == 0 {
			res = append(res, s[i])
			flag = 1
		}
	}
	return string(res)
}

func main() {
	s := "   the    sky  is blue   "
	fmt.Println(reverseWords(s))
}
