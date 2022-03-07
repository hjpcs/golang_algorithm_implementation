package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}

func main() {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	//node4 := new(ListNode)
	//node5 := new(ListNode)
	node1.Val = 1
	node1.Next = node2
	node2.Val = 1
	node2.Next = node3
	node3.Val = 2
	//node3.Next = node4
	//node4.Val = 3
	//node4.Next = node5
	//node5.Val = 3
	deleteDuplicates(node1)
}
