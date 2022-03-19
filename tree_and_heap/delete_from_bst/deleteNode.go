package delete_from_bst

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// root is a leaf node
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// root has one child node
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// root has two child node
		if root.Left != nil && root.Right != nil {
			minNode := getMin(root.Right)
			root.Val = minNode.Val
			root.Right = deleteNode(root.Right, minNode.Val)
		}
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func getMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
