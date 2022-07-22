package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func deleteDuplicates(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		for fast != nil && fast.Val == slow.Val {
			fast = fast.Next
		}
		slow.Next = fast
		slow = fast
	}

	return head
}

func main() {
	a := new(ListNode)
	a.Val = 1
	b := new(ListNode)
	b.Val = 1
	c := new(ListNode)
	c.Val = 2
	a.Next, b.Next = b, c

	d := deleteDuplicates(a)
	fmt.Println(d.Val, d.Next.Val)
}
