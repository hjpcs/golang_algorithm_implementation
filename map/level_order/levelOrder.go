package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ele := make([]int, 0)
	res := make([][]int, 0)
	ele = append(ele, root.Val)
	res = append(res, ele)
	for len(queue) != 0 {
		length := len(queue)
		tempEle := make([]int, 0)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
				tempEle = append(tempEle, cur.Left.Val)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				tempEle = append(tempEle, cur.Right.Val)
			}
		}
		if len(tempEle) != 0 {
			res = append(res, tempEle)
		}
	}
	return res
}

func main() {
	node3 := &TreeNode{Val: 3}
	node9 := &TreeNode{Val: 9}
	node20 := &TreeNode{Val: 20}
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	node3.Left = node9
	node3.Right = node20
	node20.Left = node15
	node20.Right = node7
	fmt.Println(levelOrder(node3))
}
