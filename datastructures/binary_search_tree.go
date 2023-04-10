package datastructures

// node
type node struct {
	Value int
	Left  *node
	Right *node
}

// new node
func newNode(value int) *node {
	return &node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

// binary search tree
type BinarySearchTree struct {
	Root   *node
	Length int
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		Root: nil,
	}
}

// insert
func (bst *BinarySearchTree) Insert(value int) {
	if bst.Root == nil {
		bst.Root = newNode(value)
		bst.Length++
	} else {
		bst.insertNode(bst.Root, value)
	}
}

// insert node
func (bst *BinarySearchTree) insertNode(current *node, value int) {
	if value < current.Value {
		if current.Left == nil {
			current.Left = newNode(value)
			bst.Length++
		} else {
			bst.insertNode(current.Left, value)
		}
	} else {
		if current.Right == nil {
			current.Right = newNode(value)
			bst.Length++
		} else {
			bst.insertNode(current.Right, value)
		}
	}

}

// contains
func (bst *BinarySearchTree) Contains(Root *node, value int) bool {
	if Root == nil {
		return false
	}
	if Root.Value == value {
		return true
	} else if value < Root.Value {
		return bst.Contains(Root.Left, value)
	} else {
		return bst.Contains(Root.Right, value)
	}
}

// remove
func (bst *BinarySearchTree) Remove(value int) bool {
	nodeToRemove := bst.FindNode(bst.Root, value)
	if nodeToRemove == nil {
		return false
	}
	parent := bst.FindParent(bst.Root, value)
	if bst.Length == 1 {
		bst.Root = nil
	} else if nodeToRemove.Left == nil && nodeToRemove.Right == nil {
		if nodeToRemove.Value < parent.Value {
			parent.Left = nodeToRemove.Right
		} else {
			parent.Right = nodeToRemove.Right
		}
	} else if nodeToRemove.Left != nil && nodeToRemove.Right != nil {
		next := nodeToRemove.Right
		for next.Left != nil {
			next = next.Left
		}
		if next != nodeToRemove.Right {
			bst.Remove(next.Value)
			nodeToRemove.Value = next.Value
		} else {
			nodeToRemove.Value = next.Value
			nodeToRemove.Right = nodeToRemove.Right.Right
		}
	} else {
		next := nodeToRemove.Right
		if nodeToRemove.Left != nil {
			next = nodeToRemove.Right
		} else {
			next = nodeToRemove.Left
		}
		if bst.Root == nodeToRemove {
			bst.Root = next
		} else if parent.Left == nodeToRemove {
			parent.Left = next
		} else if parent.Right == nodeToRemove {
			parent.Right = next
		}
	}
	bst.Length--
	return true
}

// findParent
func (bst *BinarySearchTree) FindParent(root *node, value int) *node {
	if value == root.Value {
		return nil
	}
	if value < root.Value {
		if root.Left == nil {
			return nil
		} else if root.Left.Value == value {
			return root
		} else {
			return bst.FindParent(root.Left, value)
		}
	} else {
		if root.Right == nil {
			return nil
		} else if root.Right.Value == value {
			return root
		} else {
			return bst.FindParent(root.Right, value)
		}
	}
}

// findNode
func (bst *BinarySearchTree) FindNode(root *node, value int) *node {
	if root == nil {
		return nil
	}
	if root.Value == value {
		return root
	} else if value < root.Value {
		return bst.FindNode(root.Left, value)
	} else {
		return bst.FindNode(root.Right, value)
	}
}

// FindMin
func (bst *BinarySearchTree) FindMin(root *node) *node {
	if root.Left == nil {
		return root
	} else {
		return bst.FindMin(root.Left)
	}
}

// Inorder
func (bst *BinarySearchTree) Inorder(root *node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*node, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Value)
		cur = node.Right
	}
	return res
}

func (bst *BinarySearchTree) Preorder(root *node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Value)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

func (bst *BinarySearchTree) Postorder(root *node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack1 := make([]*node, 0)
	stack2 := make([]int, 0)
	stack1 = append(stack1, root)
	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, node.Value)
		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}
	for i := len(stack2) - 1; i >= 0; i-- {
		res = append(res, stack2[i])
	}
	return res
}
