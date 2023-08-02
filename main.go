package main

import "fmt"

// Node represents a single node in the doubly linked list
type Node struct {
	data int   // The data stored in the node, can be of any data type
	next *Node // A pointer to the next node in the list
	prev *Node // A pointer to the previous node in the list
}

// LinkedList represents a doubly linked list
type LinkedList struct {
	head *Node // A pointer to the first node (head) of the linked list.
}

// AddNode adds a new node to the end of the linked list
func (list *LinkedList) AddNode(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	newNode.prev = current
}

// Forwards prints the elements of the linked list forwards
func (list *LinkedList) Forwards() {
	current := list.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

// Backwards prints the elements of the linked list backwards
func (list *LinkedList) Backwards() {
	current := list.head
	for current.next != nil {
		current = current.next
	}
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.prev
	}
	fmt.Println()
}

func main() {
	// Create a new linked list
	myList := LinkedList{}

	// elements to the linked list
	myList.AddNode(1)
	myList.AddNode(2)
	myList.AddNode(3)
	myList.AddNode(4)
	myList.AddNode(5)

	// Display the linked list forwards and backwards
	fmt.Println("Linked List (Forwards):")
	myList.Forwards()

	fmt.Println("Linked List (Backwards):")
	myList.Backwards()
}
