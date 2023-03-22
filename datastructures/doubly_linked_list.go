package datastructures

import "fmt"

// DoublyLinkedList is a doubly linked list
type DoublyLinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

// NewDoubleLinkedList creates a new doubly linked list
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Length: 0,
		Head:   nil,
		Tail:   nil,
	}
}

// Add adds a node to the end of the doubly linked list
func (l *DoublyLinkedList) Add(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		l.Length++
		return
	}
	n.Prev = l.Tail
	l.Tail.Next = n
	l.Tail = n
	l.Length++
}

// Remove a node from the doubly linked list
func (l *DoublyLinkedList) Remove(n *Node) bool {
	if l.Head == nil {
		return false
	}
	if l.Head.Value == n.Value {
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Head = l.Head.Next
			l.Head.Prev = nil
		}

		l.Length--
		return true
	}
	curr := l.Head.Next
	for curr != nil && curr.Value != n.Value {
		curr = curr.Next
	}
	if curr == l.Tail {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
		l.Length--
		return true
	} else if curr != nil {
		curr.Prev.Next = curr.Next
		curr.Next.Prev = curr.Prev
		l.Length--
		return true
	}
	return false
}

// Traverse the doubly linked list
func (l *DoublyLinkedList) Traverse() {
	n := l.Tail
	for n != nil {
		fmt.Println(n.Value)
		n = n.Prev
	}
}
