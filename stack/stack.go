package stack

import "github.com/GuilhermeNogueira/brackets-validator/lists"

type Stack interface {
	Push(interface{})
	Pop()
	Size()
	IsEmpty() bool
	removeLast()
	addLast()
}

type StackImpl struct {
	lists.DoublyLinkedList
}
