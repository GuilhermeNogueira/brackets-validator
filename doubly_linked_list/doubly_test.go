package doubly_linked_list

import (
	"testing"
)

func TestShouldGetCorrectSize(t *testing.T) {
	list := NewDoublyLinkedList()
	if list.size != 0 {
		t.Error("Initial size is not zero")
	}

	list.AddFirst(1)

	if list.size != 1 {
		t.Error("List size should be 1")
	}

	listSize := list.size

	i := 0
	for i <= 5 {
		list.AddFirst(i)
		i++
	}

	if s := list.size; s != listSize+i {
		t.Errorf("List size is not %v", listSize+i)
	}

	t.Log("List size is fine")
}

func TestShouldAddWhenListIsEmpty(t *testing.T) {
	list := NewDoublyLinkedList()

	list.AddFirst("abc")

	if list.size != 1 {
		t.Error("Item not added correctly")
	}

	list = NewDoublyLinkedList()

	list.AddLast("abc")

	if list.size != 1 {
		t.Error("Item not added correctly")
	}
}

func TestShouldBeEmpty(t *testing.T) {
	list := NewDoublyLinkedList()
	if s := list.IsEmpty(); s == false {
		t.Error("List should be empty")
	}
}
