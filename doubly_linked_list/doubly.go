package doubly_linked_list

import (
	"errors"
	"github.com/pedidosya/peya-go/logs"
)

type Node struct {
	element interface{}
	next    *Node
	prev    *Node
}

func NewNode(element interface{}, next *Node, prev *Node) *Node {
	return &Node{element: element, next: next, prev: prev}
}

type DoublyLinkedList interface {
	Size() int
	IsEmpty() bool
	AddFirst(*interface{})
	AddLast(*interface{})
	RemoveFirst() (*interface{}, error)
	RemoveLast() (*interface{}, error)
}

type DoublyLinkedListImpl struct {
	head *Node
	tail *Node
	size int
}

func (d *DoublyLinkedListImpl) Size() int {
	return d.size
}

func (d *DoublyLinkedListImpl) IsEmpty() bool {
	return d.size == 0
}

func (d *DoublyLinkedListImpl) AddFirst(val *interface{}) {
	tmpNode := NewNode(val, d.head, nil)
	if d.head != nil {
		d.head.prev = tmpNode
	}
	d.head = tmpNode
	if d.tail == nil {
		d.tail = tmpNode
	}
	d.size++
	logs.Infof("Element %s added", val)
}

func (d *DoublyLinkedListImpl) AddLast(val *interface{}) {
	tmpNode := NewNode(val, nil, d.tail)
	if d.tail != nil {
		d.tail.next = tmpNode
	}
	d.tail = tmpNode
	if d.head == nil {
		d.head = tmpNode
	}
	d.size++
	logs.Infof("Element %s added", val)
}

func (d *DoublyLinkedListImpl) RemoveFirst() (*interface{}, error) {
	if d.size == 0 {
		return nil, errors.New("list is empty")
	}
	tmpNode := d.head
	d.head = tmpNode.next
	d.head.prev = nil
	d.size--
	logs.Info("First item removed")
	return &tmpNode.element, nil
}

func (d *DoublyLinkedListImpl) RemoveLast() (*interface{}, error) {
	if d.size == 0 {
		return nil, errors.New("list is empty")
	}
	tmpNode := d.tail
	d.tail = tmpNode.prev
	d.tail.next = nil
	d.size--

	logs.Info("Last item removed")
	return &tmpNode.element, nil
}

func NewDoublyLinkedListImpl(head *Node, tail *Node, size int) *DoublyLinkedListImpl {
	return &DoublyLinkedListImpl{head: head, tail: tail, size: size}
}
