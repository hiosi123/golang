package linkedList

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Len  int
}

func (l *LinkedList) Add(value int) {
	newNode := new(Node)
	newNode.Value = value

	if l.Head == nil {
		l.Head = newNode
	} else {
		iterator := l.Head
		for ; iterator.Next != nil; iterator = iterator.Next {
		}
		iterator.Next = newNode
	}

}

func (l *LinkedList) Remove(value int) {
	var previous *Node

	for current := l.Head; current != nil; current = current.Next {
		if l.Head == current {
			l.Head = current.Next
		} else {
			previous.Next = current.Next
			return
		}
		previous = current
	}

}

func (l LinkedList) Get(value int) *Node {
	for iterator := l.Head; iterator != nil; iterator = iterator.Next {
		if iterator.Value == value {
			return iterator
		}
	}
	return nil
}

func (l LinkedList) String() string {
	sb := strings.Builder{}

	for iterator := l.Head; iterator != nil; iterator = iterator.Next {
		sb.WriteString(fmt.Sprintf("%d ", iterator.Value))
	}

	return sb.String()
}
