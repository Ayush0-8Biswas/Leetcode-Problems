package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isSameTree(p *TreeNode, q *TreeNode) bool {
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
	if !isSameTree(p.Left, q.Left) {
		return false
	}
	return isSameTree(p.Right, q.Right)
}

func main() {
	a1, b1, c1 := new(TreeNode), new(TreeNode), new(TreeNode)
	a1.Val, b1.Val, c1.Val = 2, 1, 1
	a1.Left, a1.Right = b1, c1

	a2, b2, c2 := new(TreeNode), new(TreeNode), new(TreeNode)
	a2.Val, b2.Val, c2.Val = 1, 1, 2
	a1.Left, a1.Right = b1, c1

	fmt.Println(isSameTree(a1, a2))
}
