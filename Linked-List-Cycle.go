package main

func hasCycle(head *ListNode) bool {
	var fast, slow = head, head

	var isCyclic = false

	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			isCyclic = true
			break
		}
	}

	return isCyclic
}