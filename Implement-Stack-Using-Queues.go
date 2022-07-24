package main

type MyStack struct {
	Val  int
	Next *MyStack
}

//func Constructor() MyStack {
//	return *new(MyStack)
//}

func (this *MyStack) Push(x int) {
	var newNode = new(MyStack)
	newNode.Val, newNode.Next = x, this.Next
	this.Next = newNode
}

func (this *MyStack) Pop() int {
	var result = this.Next.Val
	this.Next = this.Next.Next
	return result
}

func (this *MyStack) Top() int {
	return this.Next.Val
}

func (this *MyStack) Empty() bool {
	return this.Next == nil
}
