package main

import (
	"fmt"
)

// Node represents a node in the linked list
type Node struct {
	Date int
	Next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	Head *Node
}

// InsertAtBeginning inserts a new node at the beginning of the list
func (ll *LinkedList) InsertAtBeginning(date int) {
	newNode := &Node{Date: date, Next: ll.Head}
	ll.Head = newNode
}
// This method creates a new node with the given data, sets its Next pointer to the current head of the list, and then updates the list's Head to point to the new node.


// InsertAtEnd inserts a new node at the end of the list
func (ll *LinkedList) InsertAtEnd(date int) {
	newNode := &Node{Date: date, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}
// This method creates a new node with the given data and a nil Next pointer. If the list is empty, it sets the Head to the new node. Otherwise, it traverses the list to find the last node and sets its Next pointer to the new node.


// DeleteWithValue deletes the first node with the give value
func (ll *LinkedList) DeleteWithValue(date int) {
	if ll.Head == nil {
		return // List is empty
	}

	if ll.Head.Date == date {
		ll.Head = ll.Head.Next //Delete the head node
		return
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Date == date {
			current.Next = current.Next.Next // Bypass the node to
			return // delete it
		}
		current = current.Next
	}
}
// This method first checks if the list is empty. If not, it checks if the head node contains the value to be deleted. If so, it updates the Head to point to the next node. Otherwise, it traverses the list to find the node to be deleted and updates the Next pointer of the previous node to skip the node to be deleted.

// Display prints the elements of the linked list
func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Date)
		current = current.Next
	}
	fmt.Println("nil")
}
// This method traverses the list from the Head to the end, printing the data of each node.

func main() {
	ll := LinkedList{}
	ll.InsertAtEnd(10)
	ll.InsertAtEnd(20)
	ll.InsertAtBeginning(5)
	ll.Display() // Output: 5 -> 10 -> 20 -> nil

	ll.DeleteWithValue(10)
	ll.Display() // Output: 5 -> 20 -> nil

	ll.DeleteWithValue(5)
	ll.Display() // Output: 20 -> nil

	ll.DeleteWithValue(20)
	ll.Display() // Output: nil
}