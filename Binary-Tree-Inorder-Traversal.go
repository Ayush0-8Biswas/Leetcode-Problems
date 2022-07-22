package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}

func main() {
	a := new(TreeNode)
	a.Val = 1
	b := new(TreeNode)
	b.Val = 2
	c := new(TreeNode)
	c.Val = 3
	a.Right = b
	b.Left = c

	fmt.Println(inorderTraversal(a))
}
