package main

import "fmt"

func Constructor() MyLinkedList {
	return MyLinkedList{nil, 0}
}

type Node struct {
	Value int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	Capacity int 
}

func (l *MyLinkedList) AddAtHead(value int) {
	newNode := &Node{
		Value: value,
		Next: l.Head,
	}

	l.Head = newNode
	l.Capacity++
}

func (l *MyLinkedList) AddAtTail(value int) {
	newNode := &Node{
		Value: value,
	}
	current := l.Head
	if l.Head == nil {
		l.Head = newNode
		l.Capacity++
		return 
	}
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	l.Capacity++
}

func (l *MyLinkedList) AddAtIndex(index, value int) {
	if index < 0 || index > l.Capacity {
		return
	}
	if index == 0 {
		l.AddAtHead(value)
		return
	}
	curr := l.Head
	for i := 1; i < index; i++ {
		curr = curr.Next
	}
	newNode := &Node{Value: value}
	if curr == nil {
		curr = newNode
		return
	}

	newNode.Next = curr.Next
	curr.Next = newNode
	l.Capacity++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.Capacity {
		return
	}
	if index == 0 && l.Head != nil {
		l.Head = l.Head.Next
		l.Capacity--
		return
	}
	curr := l.Head
	for i := 1; i < index; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	l.Capacity--
}

func (l *MyLinkedList) Get(index int) int {
	if index >= l.Capacity || index < 0 { 
		return -1
	}
	curNode := l.Head
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}
	return curNode.Value
}

func (l *MyLinkedList) printList() {
    current := l.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Value)
        current = current.Next
    }
    fmt.Println("NULL")
}
func main() {
	var l *MyLinkedList	

	l = &MyLinkedList{}
	l.AddAtHead(7)
	l.AddAtHead(2)
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(3, 0)
	l.DeleteAtIndex(2)
	l.AddAtHead(6)
	l.AddAtTail(4)
	l.AddAtHead(4)
	l.AddAtIndex(5,0)
	l.AddAtHead(6)
	l.AddAtHead(1)

	l.printList()
}