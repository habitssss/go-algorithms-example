package datastructures

import (
	"testing"
)

func TestAdd(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.Add(NewNode(i))
	}
	if list.Length != 10 {
		t.Errorf("Expected length 10, got %d", list.Length)
	}
	if list.Head.Value != 0 {
		t.Errorf("Expected head value 0, got %d", list.Head.Value)
	}
	if list.Tail.Value != 9 {
		t.Errorf("Expected tail value 9, got %d", list.Tail.Value)
	}
}

func TestContains(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.Add(NewNode(i))
	}
	if !list.Contains(NewNode(0)) {
		t.Errorf("Expected true, got false")
	}
	if list.Contains(NewNode(100)) {
		t.Errorf("Expected false, got true")
	}
}

func TestRemove(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.Add(NewNode(i))
	}
	if list.Length != 10 {
		t.Errorf("Expected length 10, got %d", list.Length)
	}
	if list.Head.Value != 0 {
		t.Errorf("Expected head value 0, got %d", list.Head.Value)
	}
	if list.Tail.Value != 9 {
		t.Errorf("Expected tail value 9, got %d", list.Tail.Value)
	}
	result := list.Remove(NewNode(0))
	if !result {
		t.Errorf("Expected true, got false")
	}
	if list.Length != 9 {
		t.Errorf("Expected length 9, got %d", list.Length)
	}
	if list.Head.Value != 1 {
		t.Errorf("Expected head value 1, got %d", list.Head.Value)
	}
}

func TestTraverse(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.Add(NewNode(i))
	}
	list.Print()
	if list.Head.Value != 0 {
		t.Errorf("Expected head value 0, got %d", list.Head.Value)
	}
	if list.Tail.Value != 9 {
		t.Errorf("Expected tail value 9, got %d", list.Tail.Value)
	}
}

func TestReverse(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.Add(NewNode(i))
	}
	list.Reverse()
	list.Print()
	if list.Head.Value != 9 {
		t.Errorf("Expected head value 9, got %d", list.Head.Value)
	}
	if list.Tail.Value != 0 {
		t.Errorf("Expected tail value 0, got %d", list.Tail.Value)
	}
}
