package datastructures

import "testing"

// test for doubly linked list add
func TestDoublyLinkedListAdd(t *testing.T) {
	list := NewDoublyLinkedList()
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

// test for doubly linked list remove
func TestDoublyLinkedListRemove(t *testing.T) {
	list := NewDoublyLinkedList()
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
	result := list.Remove(NewNode(7))
	if !result {
		t.Errorf("Expected true, got false")
	}
	if list.Length != 9 {
		t.Errorf("Expected length 9, got %d", list.Length)
	}
}

func TestDoublyLinkedListTraverse(t *testing.T) {
	list := NewDoublyLinkedList()
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
	list.Traverse()
}
