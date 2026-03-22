package linkedlist

type LinkedList struct {
	head *ListNode
	tail *ListNode
	size int64
}

type ListNode struct {
	Data any
	Next *ListNode
	Prev *ListNode
}
