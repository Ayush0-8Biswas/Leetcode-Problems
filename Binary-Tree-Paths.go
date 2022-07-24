package main

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	var leftData, rightData = binaryTreePaths(root.Left), binaryTreePaths(root.Right)
	var result = make([]string, 0, len(leftData)+len(rightData))
	var x = strconv.Itoa(root.Val) + "->"

	for _, v := range leftData {
		result = append(result, x+v)
	}
	for _, v := range rightData {
		result = append(result, x+v)
	}

	return result
}
