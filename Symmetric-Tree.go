package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	return isMirrorTree(root.Right, root.Left)
}

func isMirrorTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil {
		return false
	} else if q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if !isMirrorTree(p.Left, q.Left) {
		return false
	}
	return isMirrorTree(p.Right, q.Right)
}

func main() {

}
