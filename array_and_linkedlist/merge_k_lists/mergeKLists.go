package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode
	for i := 0; i < len(lists); i++ {
		result = mergeTwoLists(result, lists[i])
	}
	return result
}

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	head := new(ListNode)
	tail, aPtr, bPtr := head, a, b
	for aPtr != nil && bPtr != nil {
		if aPtr.Val < bPtr.Val {
			tail.Next = aPtr
			aPtr = aPtr.Next
		} else {
			tail.Next = bPtr
			bPtr = bPtr.Next
		}
		tail = tail.Next
	}
	if aPtr != nil {
		tail.Next = aPtr
	} else {
		tail.Next = bPtr
	}
	return head.Next
}

func main() {
	a1 := new(ListNode)
	a2 := new(ListNode)
	a3 := new(ListNode)
	a1.Val = 1
	a1.Next = a2
	a2.Val = 4
	a2.Next = a3
	a3.Val = 5

	b1 := new(ListNode)
	b2 := new(ListNode)
	b3 := new(ListNode)
	b1.Val = 1
	b1.Next = b2
	b2.Val = 3
	b2.Next = b3
	b3.Val = 4

	c1 := new(ListNode)
	c2 := new(ListNode)
	c1.Val = 2
	c1.Next = c2
	c2.Val = 6

	var lists []*ListNode
	lists = append(lists, a1, b1, c1)

	result := mergeKLists(lists)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
