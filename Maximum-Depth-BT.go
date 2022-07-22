package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1 + max(maxDepth(root.Left), maxDepth(root.Right))
	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
