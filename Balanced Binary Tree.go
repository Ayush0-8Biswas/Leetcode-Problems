package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if mod(maxDepth(root.Left)-maxDepth(root.Right)) > 1 {
		return false
	}

	if !isBalanced(root.Right) {
		return false
	}

	return isBalanced(root.Left)
}

func mod(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	depth := 1 + max(maxDepth(root.Left), maxDepth(root.Right))
//	return depth
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
