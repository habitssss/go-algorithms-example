package datastructures

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}
