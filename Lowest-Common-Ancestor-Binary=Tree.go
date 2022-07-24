package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var a, b, c = root.Val, p.Val, q.Val

	if (a >= b && a <= c) || (a >= c && a <= b) {
		return root
	}

	if a > b && a > c {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return lowestCommonAncestor(root.Right, p, q)
}
