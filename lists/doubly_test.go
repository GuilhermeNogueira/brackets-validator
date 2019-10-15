package lists

import (
	"testing"
)

func TestShouldGetCorrectSize(t *testing.T) {
	list := NewDoublyLinkedList()
	if list.Size() != 0 {
		t.Error("Initial size is not zero")
	}

	list.AddFirst(1)

	if list.Size() != 1 {
		t.Error("List size should be 1")
	}

	listSize := list.Size()

	i := 0
	for i <= 5 {
		list.AddFirst(i)
		i++
	}

	if s := list.Size(); s != listSize+i {
		t.Errorf("List size is not %v", listSize+i)
	}

	t.Log("List size is fine")
}

func TestShouldAddWhenListIsEmpty(t *testing.T) {
	list := NewDoublyLinkedList()

	list.AddFirst("abc")

	if list.Size() != 1 {
		t.Error("Item not added correctly")
	}

	list = NewDoublyLinkedList()

	list.AddLast("abc")

	if list.Size() != 1 {
		t.Error("Item not added correctly")
	}
}

func TestShouldBeEmpty(t *testing.T) {
	list := NewDoublyLinkedList()
	if s := list.IsEmpty(); s == false {
		t.Error("List should be empty")
	}
}

func TestDoublyLinkedListImpl_IndexOf(t *testing.T) {
	list := newDefaultList()
	a := "A"
	b := "B"
	c := "C"

	if list.IndexOf(a) != 0 {
		t.Error("Index should be 0 ")
	}

	if list.IndexOf(b) != 1 {
		t.Error("Index should be 1 ")
	}
	if list.IndexOf(c) != 2 {
		t.Error("Index should be 2 ")
	}

	if list.IndexOf("z") != -1 {
		t.Error("Index should be -1 ")
	}
}

func TestDoublyLinkedListImpl_Insert(t *testing.T) {
	list := newDefaultList()

	item := "ITEM"

	list.Insert(item, 1)

	index := list.IndexOf(item)

	if index != 1 {
		t.Error("Index should be 1")
	}
}

func TestDoublyLinkedListImpl_String(t *testing.T) {

	list := newDefaultList()

	s := list.String()

	expected := "A -> B -> C"

	if s != expected {
		t.Errorf("%v should be %v", s, expected)
	}
}

func TestDoublyLinkedListImpl_RemoveAt(t *testing.T) {
	list := newDefaultList()
	at, _ := list.RemoveAt(1)

	if *at != "B" {
		t.Error("Remove fail")
	}

	if list.Size() != 2 {
		t.Error("List sie should be 2")
	}
}

func newDefaultList() DoublyLinkedList {
	list := NewDoublyLinkedList()
	list.AddLast("A")
	list.AddLast("B")
	list.AddLast("C")
	return list
}
