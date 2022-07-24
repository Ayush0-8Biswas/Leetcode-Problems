package main

type MyQueue struct {
	last  *QueueNode
	first *QueueNode
}

type QueueNode struct {
	val  int
	next *QueueNode
}

//func Constructor() MyQueue {
//	return *new(MyQueue)
//}

func (this *MyQueue) Push(x int) {
	var newNode = new(QueueNode)
	newNode.val = x
	if this.last == nil {
		this.first, this.last = newNode, newNode
	} else {

		this.last.next = newNode
		this.last = newNode
	}
}

func (this *MyQueue) Pop() int {
	var result = this.first.val
	this.first = this.first.next

	if this.first == nil {
		this.last = nil
	}

	return result
}

func (this *MyQueue) Peek() int {
	return this.first.val
}

func (this *MyQueue) Empty() bool {
	return this.first == nil
}
