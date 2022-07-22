package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func sortedArrayToBST(nums []int) *TreeNode {
	size := len(nums)
	if size == 0 {
		return nil
	}
	mid := (size - 1) / 2
	root := new(TreeNode)
	root.Val = nums[mid]
	root.Right = sortedArrayToBST(nums[mid+1:])
	root.Left = sortedArrayToBST(nums[:mid])

	return root
}

func main() {
	sortedArrayToBST([]int{-10, -3, 0, 5, 9})
}
