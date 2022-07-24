package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func reverseList(head *ListNode) *ListNode {
//	var previous, current, next *ListNode = nil, head, nil
//
//	for current != nil {
//		next = current.Next
//		current.Next = previous
//		previous = current
//		current = next
//	}
//
//	return previous
//}
//
//func isPalindrome(head *ListNode) bool {
//	var slow, fast = head, head
//
//	for fast != nil && fast.Next != nil {
//		slow, fast = slow.Next, fast.Next.Next
//	}
//	if fast != nil {
//		slow, fast = slow.Next, head
//	} else {
//		fast = head
//	}
//
//	slow = reverseList(slow)
//
//	for slow != nil {
//		if slow.Val != fast.Val {
//			return false
//		}
//		slow, fast = slow.Next, fast.Next
//	}
//
//	return true
//}

func main() {
	var a, b, c, d = new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	a.Val, b.Val, c.Val, d.Val = 1, 2, 2, 1
	a.Next, b.Next, c.Next = b, c, d
	fmt.Println(isPalindrome(a))
}
