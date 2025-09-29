package main

import "fmt"

type Node struct {
	data int64
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertAtEnd(data int64) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (l *LinkedList) InsertAtBegin(data int64) {
	newNode := &Node{data: data}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) Delete(data int64) {
	if l.head == nil {
		return
	}

	if l.head.data == data {
		l.head = l.head.next
		return
	}

	current := l.head

	for current.next != nil && current.next.data != data {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}

func (l *LinkedList) Traverse() {
	if l.head == nil {
		fmt.Println("No Data")
		return
	}

	current := l.head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}

	fmt.Println("Done")
}

func main() {
	list := &LinkedList{}
	list.InsertAtEnd(3)
	list.InsertAtEnd(6)
	list.InsertAtEnd(9)
	list.InsertAtBegin(100)
	list.InsertAtEnd(12)
	list.InsertAtEnd(64)
	list.InsertAtEnd(245)
	list.Traverse()
	list.Delete(9)
	list.Traverse()


}