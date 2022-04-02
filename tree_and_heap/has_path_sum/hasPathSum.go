package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
	node1 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 8}
	node4 := &TreeNode{Val: 11}
	node5 := &TreeNode{Val: 13}
	node6 := &TreeNode{Val: 4}
	node7 := &TreeNode{Val: 7}
	node8 := &TreeNode{Val: 2}
	node9 := &TreeNode{Val: 1}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node3.Left = node5
	node3.Right = node6
	node4.Left = node7
	node4.Right = node8
	node6.Right = node9
	fmt.Println(hasPathSum(node1, 22))
}
