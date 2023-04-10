package datastructures

import (
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(13)
	tree.Insert(14)
	tree.Insert(15)
	tree.Insert(16)
	tree.Insert(17)
	tree.Insert(18)

	if tree.Length != 18 {
		t.Errorf("Expected length 18, got %d", tree.Length)
	}
}

// contains
func TestBinarySearchTreeContains(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(13)
	tree.Insert(14)
	tree.Insert(15)
	tree.Insert(16)
	tree.Insert(17)
	tree.Insert(18)

	if !tree.Contains(tree.Root, 3) {
		t.Errorf("Expected true, got false")
	}
	if !tree.Contains(tree.Root, 19) {
		t.Errorf("Expected false, got true")
	}
}

func TestBinarySearchTreeRemove(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(13)
	tree.Insert(14)
	tree.Insert(15)
	tree.Insert(16)
	tree.Insert(17)
	tree.Insert(18)

	if tree.Length != 18 {
		t.Errorf("Expected length 18, got %d", tree.Length)
	}
	if !tree.Contains(tree.Root, 3) {
		t.Errorf("Expected false, got true")
	}
	if tree.Contains(tree.Root, 19) {
		t.Errorf("Expected false, got true")
	}
	result := tree.Remove(3)
	if !result {
		t.Errorf("Expected true, got false")
	}
	if tree.Length != 17 {
		t.Errorf("Expected length 17, got %d", tree.Length)
	}
}

func TestBinarySearchTreeFindMin(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(13)
	tree.Insert(14)
	tree.Insert(15)
	tree.Insert(16)
	tree.Insert(17)
	tree.Insert(18)

	if tree.FindMin(tree.Root).Value != 1 {
		t.Errorf("Expected 1, got %d", tree.FindMin(tree.Root).Value)
	}
}

func TestBinarySearchTreeInorder(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(13)
	tree.Insert(14)
	tree.Insert(15)
	tree.Insert(16)
	tree.Insert(17)
	tree.Insert(18)
	result := tree.Inorder(tree.Root)
	t.Log(result)
}

func TestBinarySearchTreePreorder(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(13)
	tree.Insert(14)
	tree.Insert(15)
	tree.Insert(16)
	tree.Insert(17)
	tree.Insert(18)
	result := tree.Preorder(tree.Root)
	t.Log(result)
}

func TestBinarySearchTreePostorder(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(13)
	tree.Insert(14)
	tree.Insert(15)
	tree.Insert(16)
	tree.Insert(17)
	tree.Insert(18)
	result := tree.Postorder(tree.Root)
	t.Log(result)
}
