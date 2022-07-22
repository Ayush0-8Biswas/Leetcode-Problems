package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1 + specialMin(minDepth(root.Left), minDepth(root.Right))
	return depth
}

func specialMin(a, b int) int {
	if b == 0 {
		return a
	}
	if a == 0 {
		return b
	}
	if a < b {
		return a
	}
	return b
}

func main() {

}
