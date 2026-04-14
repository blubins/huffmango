package linkedlist

import (
	"fmt"
	"strings"
)

// returns a pointer to a new linked list
func New() *LinkedList {
	return &LinkedList{}
}

// returns the size of the linked list
func (l *LinkedList) Size() int64 {
	return l.size
}

// returns string representation of the linked list inorder
func (l *LinkedList) String() string {
	var out strings.Builder
	iterator := l.head
	for iterator != nil {
		out.Write(fmt.Appendf(nil, "%v", iterator.Data))
		out.Write([]byte(","))
		iterator = iterator.Next
	}
	result := strings.TrimSuffix(out.String(), ",")
	return result
}

// prints the linked list to stdout
func (l *LinkedList) Print() {
	var out strings.Builder
	iterator := l.head
	out.Write([]byte("Linked List: {"))
	for iterator != nil {
		out.Write(fmt.Appendf(nil, "%v", iterator.Data))
		out.Write([]byte(", "))
		iterator = iterator.Next
	}
	result := strings.TrimSuffix(out.String(), ", ")
	fmt.Println(result + "}")
}

// inserts into the linked list sorted based off filter function parameter
func (l *LinkedList) InsertSorted(data any, less func(a, b any) bool) {
	newNode := &ListNode{Data: data}

	// case for new head
	if l.head == nil || less(data, l.head.Data) {
		newNode.Next = l.head
		if l.head != nil {
			l.head.Prev = newNode
		} else {
			l.head = newNode
		}
		l.head = newNode
		l.size++
		return
	}
	// generic case, middle of list somewhere
	cur := l.head
	for cur.Next != nil && !less(data, cur.Next.Data) {
		cur = cur.Next
	}

	// case for new tail
	newNode.Next = cur.Next
	newNode.Prev = cur
	if cur.Next != nil {
		cur.Next.Prev = newNode
	} else {
		l.tail = newNode
	}
	cur.Next = newNode
	l.size++
}

// appends an item to the end of the linked list
func (l *LinkedList) Append(data any) {
	newNode := ListNode{
		Data: data,
		Next: nil,
		Prev: l.tail,
	}

	if l.tail == nil {
		l.head = &newNode
	} else {
		l.tail.Next = &newNode
	}

	l.tail = &newNode
	l.size++
}

// deletes an item at a specified index of the linked list
func (l *LinkedList) DeleteNode(pos int) *ListNode {
	// check if invalid index or list is empty
	if pos < 1 || l.head == nil {
		return nil
	}
	// start at the head
	tmp := l.head
	// go to the node we want to delete with null checking
	for i := 1; i < pos && tmp != nil; i++ {
		tmp = tmp.Next
	}
	// make sure its not 1 off and still null
	if tmp == nil {
		return nil
	}
	// assign prev node next to next
	if tmp.Prev != nil {
		tmp.Prev.Next = tmp.Next
	} else {
		// case for new head
		l.head = tmp.Next
	}

	// case for new tail
	if tmp.Next != nil {
		tmp.Next.Prev = tmp.Prev
	} else {
		l.tail = tmp.Prev
	}

	l.size--
	return tmp
}
