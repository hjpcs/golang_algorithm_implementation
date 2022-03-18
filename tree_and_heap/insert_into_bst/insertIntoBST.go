package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		//root = new(TreeNode)
		//root.Val = val
		root = &TreeNode{Val: val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func main() {
	node4 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2}
	node1 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 3}
	node7 := &TreeNode{Val: 7}
	node4.Left = node2
	node4.Right = node7
	node2.Left = node1
	node2.Right = node3
	fmt.Println(insertIntoBST(node4, 5))
}
