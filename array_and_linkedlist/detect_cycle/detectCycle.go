package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	meet := head
	slow = slow.Next
	for slow != meet {
		slow = slow.Next
		meet = meet.Next
	}
	return meet
}

func main() {
	node1 := new(ListNode)
	node1.Val = 3
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 0
	node4 := new(ListNode)
	node4.Val = -4
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	meet := detectCycle(node1)
	fmt.Println(meet)
}
