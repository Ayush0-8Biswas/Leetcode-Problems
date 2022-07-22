package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lengthA, lengthB = lengthList(headA), lengthList(headB)

	if lengthA > lengthB {
		for i := lengthB; i < lengthA; i++ {
			headA = headA.Next
		}
	} else if lengthB > lengthA {
		for i := lengthA; i < lengthB; i++ {
			headB = headB.Next
		}
	}

	for headA != nil || headB != nil {
		if headA == headB {
			return headA
		}
		headA, headB = headA.Next, headB.Next
	}

	return nil
}

func lengthList(head *ListNode) int {
	var result int
	for head != nil {
		head = head.Next
		result += 1
	}
	return result
}