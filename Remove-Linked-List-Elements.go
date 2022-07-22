package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeElements(head *ListNode, val int) *ListNode {
	var dummyHead = new(ListNode)
	dummyHead.Next = head
	var iterator = dummyHead

	for iterator != nil && iterator.Next != nil {
		for iterator.Next != nil && iterator.Next.Val == val {
			iterator.Next = iterator.Next.Next
		}
		iterator = iterator.Next
	}

	return dummyHead.Next
}

func main() {
	var a, b, c, d = new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	a.Val, b.Val, c.Val, d.Val = 7, 7, 7, 7
	a.Next, b.Next, c.Next = b, c, d
	a = removeElements(a, 7)
	for a != nil {
		fmt.Println(a.Val, " ")
		a = a.Next
	}
}
