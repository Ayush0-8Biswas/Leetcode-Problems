package main

import "fmt"

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res = []int{root.Val}
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)

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
