package main

import "fmt"

type Item interface{}

type ItemStack struct {
	Items []Item
	N     int
}

func (stack *ItemStack) Init(n int) *ItemStack {
	stack.Items = []Item{}
	stack.N = n
	return stack
}

func (stack *ItemStack) Push(item Item) {
	if len(stack.Items) == stack.N {
		fmt.Println("stack is full")
		return
	}
	stack.Items = append(stack.Items, item)
}

func (stack *ItemStack) Pop() Item {
	if len(stack.Items) == 0 {
		fmt.Println("stack is empty")
		return nil
	}
	item := stack.Items[len(stack.Items)-1]
	stack.Items = stack.Items[0 : len(stack.Items)-1]
	return item
}

func isValid(s string) bool {
	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if len(stack) > 0 && brackets[char] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	var stack1 ItemStack
	fmt.Println(stack1)
	stack1.Init(10)
	fmt.Println(stack1)
	stack1.Push(12345)
	fmt.Println(stack1)
	stack1.Push("hello")
	fmt.Println(stack1)
	stack1.Pop()
	fmt.Println(stack1)

	s := "()[]{}"
	fmt.Println(isValid(s))
}
