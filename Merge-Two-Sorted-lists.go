package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var res, current *ListNode = nil, nil

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			if current == nil {
				res, current = list1, list1
			} else {
				current.Next = list1
				current = list1
			}
			list1 = list1.Next
		} else {
			if current == nil {
				res, current = list2, list2
			} else {
				current.Next = list2
				current = list2
			}
			list2 = list2.Next
		}
	}

	if list1 != nil {
		if current == nil {
			return list1
		}
		current.Next = list1
	} else if list2 != nil {
		if current == nil {
			return list2
		}
		current.Next = list2
	}

	return res
}

func main() {

}
