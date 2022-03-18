package is_valid_bst

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTWithParameter(root, nil, nil)
}

func isValidBSTWithParameter(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isValidBSTWithParameter(root.Left, min, root) && isValidBSTWithParameter(root.Right, root, max)
}
