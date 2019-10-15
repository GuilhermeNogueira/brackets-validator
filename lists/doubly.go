package lists

import (
	"errors"
	"fmt"
	"github.com/pedidosya/peya-go/logs"
)

type Node struct {
	element *interface{}
	next    *Node
	prev    *Node
}

func NewNode(element *interface{}, next *Node, prev *Node) *Node {
	return &Node{element: element, next: next, prev: prev}
}

//DoublyLinkedList
/*
	TODO:
		- Add concurrency support
		- Change interface to generic

*/
type DoublyLinkedList interface {
	Size() int
	IsEmpty() bool
	AddFirst(interface{})
	AddLast(interface{})
	RemoveFirst() (*interface{}, error)
	RemoveLast() (*interface{}, error)
	String() string
	Insert(interface{}, int) bool
	RemoveAt(int) (*interface{}, error)
	IndexOf(interface{}) int
	Head() *interface{}
	Tail() *interface{}
}

type DoublyLinkedListImpl struct {
	head *Node
	tail *Node
	size int
}

func NewDoublyLinkedList() DoublyLinkedList {
	return &DoublyLinkedListImpl{}
}

func (d *DoublyLinkedListImpl) getNodeAt(pos int) (*Node, error) {

	//If wanted position is greater than our size, return nil
	if pos < 0 || pos >= d.size {
		return nil, fmt.Errorf("index out of bounds")
	}

	node := Node{}

	if pos < d.size/2 {
		node = *d.head
		for i := 0; i != pos; i++ {
			node = *node.next
		}
	} else {
		node = *d.tail
		for i := d.size - 1; i != pos; i-- {
			node = *node.prev
		}
	}

	return &node, nil
}

func (d *DoublyLinkedListImpl) remove(node Node) *interface{} {
	if node.prev == nil {
		elem, _ := d.RemoveFirst()
		return elem
	}

	if node.next == nil {
		elem, _ := d.RemoveLast()
		return elem
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	d.size--

	node.next = nil
	node.prev = nil

	return node.element
}

func (d *DoublyLinkedListImpl) String() string {
	return str(d.head, nil)
}

func str(node *Node, val *string) string {
	if node == nil {
		return *val
	}

	nextVal := fmt.Sprintf("%v", *node.element)

	if val != nil {
		nextVal = fmt.Sprintf("%v -> %v", *val, nextVal)
	}

	return str(node.next, &nextVal)
}

func (d *DoublyLinkedListImpl) Insert(val interface{}, pos int) bool {
	if pos > d.size-1 {
		return false
	}

	node, e := d.getNodeAt(pos)

	if e != nil {
		return false
	}

	newNode := NewNode(&val, node, node.prev)

	node.prev.next = newNode
	node.prev = newNode

	return true
}

func (d *DoublyLinkedListImpl) RemoveAt(pos int) (*interface{}, error) {
	node, e := d.getNodeAt(pos)

	if e != nil {
		return nil, e
	}

	remove := d.remove(*node)

	return remove, nil
}

func (d *DoublyLinkedListImpl) IndexOf(elem interface{}) int {
	node := d.head
	index := 0
	for node != nil {
		if *node.element == elem {
			return index
		}
		index++
		node = node.next
	}
	return -1
}

func (d *DoublyLinkedListImpl) Head() *interface{} {
	return d.head.element
}

func (d *DoublyLinkedListImpl) Tail() *interface{} {
	return d.tail.element
}

func (d *DoublyLinkedListImpl) Size() int {
	return d.size
}

func (d *DoublyLinkedListImpl) IsEmpty() bool {
	return d.size == 0
}

func (d *DoublyLinkedListImpl) AddFirst(val interface{}) {
	tmpNode := NewNode(&val, d.head, nil)
	if d.head != nil {
		d.head.prev = tmpNode
	}
	d.head = tmpNode
	if d.tail == nil {
		d.tail = tmpNode
	}
	d.size++
	logs.Infof("Element %v added", val)
}

func (d *DoublyLinkedListImpl) AddLast(val interface{}) {
	tmpNode := NewNode(&val, nil, d.tail)
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
	return tmpNode.element, nil
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
	return tmpNode.element, nil
}
