package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	if hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val)
}

func main() {

}
