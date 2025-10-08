package data_structures

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

// append new node to linkedList
func (l *LinkedList) Append(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
		return
	}

	lastnode := l.head
	for lastnode.next != nil {
		lastnode = lastnode.next
	}
	lastnode.next = newNode

}

// PrintList print the all element of the linkedList
func (l *LinkedList) PrintList() {
	currentNode := l.head
	for currentNode.next != nil {
		fmt.Printf("%d ->", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("nil")
}
