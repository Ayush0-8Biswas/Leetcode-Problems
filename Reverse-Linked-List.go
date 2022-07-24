package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseList(head *ListNode) *ListNode {
	var previous, current, next *ListNode = nil, head, nil

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}

func createNewNode(val int) *ListNode {
	var newNode = new(ListNode)
	newNode.Val = val
	return newNode
}

func main() {
	var a = createNewNode(1)
	a.Next = createNewNode(2)
	a.Next.Next = createNewNode(3)
	a.Next.Next.Next = createNewNode(4)
	a.Next.Next.Next.Next = createNewNode(5)

	a = reverseList(a)

	for a != nil {
		fmt.Println(a.Val)
	}
}
