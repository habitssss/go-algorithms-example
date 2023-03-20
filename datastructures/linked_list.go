package datastructures

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Length: 0,
		Head:   nil,
		Tail:   nil,
	}
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

func (l *LinkedList) Add(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
	l.Length++
}

func (l *LinkedList) Search(n *Node) *Node {
	if l.Head == nil {
		return nil
	}
	if l.Head.Value == n.Value {
		return l.Head
	}
	curr := l.Head
	for curr.Next != nil {
		if curr.Next.Value == n.Value {
			return curr.Next
		}
		curr = curr.Next
	}
	return nil
}

func (l *LinkedList) Remove(n *Node) bool {
	if l.Head == nil {
		return false
	}
	if l.Head.Value == n.Value {
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Head = l.Head.Next
		}
		l.Length--
		return true
	}
	curr := l.Head
	for curr.Next != nil {
		if curr.Next.Value == n.Value {
			curr.Next = curr.Next.Next
			l.Length--
			return true
		}
		curr = curr.Next
	}
	return false
}

func (l *LinkedList) Print() {
	if l.Head == nil {
		return
	}
	curr := l.Head
	for curr != nil {
		println(curr.Value)
		curr = curr.Next
	}
}

func (l *LinkedList) Reverse() {
	if l.Head == nil {
		return
	}
	curr := l.Head
	l.Tail = l.Head
	var prev *Node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

func (l *LinkedList) ReverseRecursive(curr *Node) *Node {
	if curr == nil {
		return nil
	}
	next := curr.Next
	curr.Next = l.ReverseRecursive(next)
	return curr
}

func (l *LinkedList) ReverseIterative() {
	if l.Head == nil {
		return
	}
	var prev *Node
	curr := l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

func (l *LinkedList) Contains(n *Node) bool {
	if l.Head == nil {
		return false
	}
	curr := l.Head
	for curr != nil {
		if curr.Value == n.Value {
			return true
		}
		curr = curr.Next
	}
	return false
}
