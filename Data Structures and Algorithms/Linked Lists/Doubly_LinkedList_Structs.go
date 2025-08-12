package main

import "fmt"

// Node represents a node in the doubly linked list
type Node struct {
	Data int
	Next *Node
	Prev *Node
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
	Head *Node
	Tail *Node //Keep track of the tail for easier insertion at the end
}


// InsertAtBeginning inserts a new node at the beginning of the list
func (dll *DoublyLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{Data: data, Next: dll.Head, Prev: nil}
	
	if dll.Head == nil {
		dll.Tail = newNode // If the list is empty, set tail to the new node
	} else {
		dll.Head.Prev = newNode
		dll.Head = newNode
	}
}

// InsertAtEnd inserts a new node at the end of the list
func (dll *DoublyLinkedList) InsertAtEnd(data int) {
	newNode := &Node{Data: data, Next: nil, Prev: dll.Tail}

	if dll.Tail == nil {
		dll.Head = newNode
		dll.Tail = newNode 
	} else {
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}
}


// DeleteWithValue deletes the first node with the given value
func (dll *DoublyLinkedList) DeleteWithValue(data int) {
	current := dll.Head

	for current != nil {
		if current.Data == data {
			if current.Prev != nil {
				current.Prev.Next = current.Next // Bypass the node to
			} else {
				dll.Head = current.Next
			}

			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				dll.Tail = current.Prev // Deleting the tail
			}
			return 
		}
		current = current.Next
	}
}

// Display prints the elements of the doubly linked list
func (dll *DoublyLinkedList) Display() {
	fmt.Print("Forward: ")
	current := dll.Head
	for current != nil {
		fmt.Printf("%d <-> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")

	fmt.Print("Reverse: ")
	current = dll.Tail
	for current != nil {
		fmt.Printf("%d <-> ", current.Data)
		current = current.Prev
	}
	fmt.Println("nil")
}

func main() {
	dll := DoublyLinkedList{}
	dll.InsertAtEnd(10)
	dll.InsertAtEnd(20)
	dll.InsertAtBeginning(5)
	dll.Display()

	dll.DeleteWithValue(10)
	dll.Display()

	dll.DeleteWithValue(5)
	dll.Display()	

	dll.DeleteWithValue(20)
	dll.Display()
}
