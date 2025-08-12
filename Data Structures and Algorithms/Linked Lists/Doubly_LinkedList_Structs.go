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
// This method creates a new node, sets its Next pointer to the current head, and its Prev pointer to nil. If the list is empty, it sets both Head and Tail to the new node. Otherwise, it updates the Prev pointer of the current head to point to the new node and updates the Head to point to the new node.


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
// This method creates a new node, sets its Next pointer to nil, and its Prev pointer to the current tail. If the list is empty, it sets both Head and Tail to the new node. Otherwise, it updates the Next pointer of the current tail to point to the new node and updates the Tail to point to the new node. This is more efficient than traversing the list to find the end, as we maintain a Tail pointer.


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
// This method traverses the list to find the node to be deleted. If the node is found, it updates the Next pointer of the previous node and the Prev pointer of the next node to bypass the node to be deleted. It also handles the cases where the head or tail is being deleted, updating the Head or Tail pointers accordingly.


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
// This method traverses the list from the Head to the end, printing the data of each node. Then, it traverses the list from the Tail to the beginning, printing the data of each node in reverse order.


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
