package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
    data int					//The data stored in the node, can be of any data type (e.g., integer, string, custom struct, etc.)
    next *Node					//A pointer to the next node in the list. In Go, pointers are represented using the * symbol.
}

// LinkedList represents a singly linked list
type LinkedList struct {
    head *Node					//A pointer to the first node (head) of the linked list.
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
}

// Display prints the elements of the linked list
func (list *LinkedList) Display() {
    current := list.head
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Println()
}

func main() {
    // Create a new linked list
    myList := LinkedList{}

    // Add elements to the linked list
    myList.AddNode(1)
    myList.AddNode(2)
    myList.AddNode(3)
    myList.AddNode(4)
	myList.AddNode(5)

    // Display the linked list
    fmt.Println("Linked List:")
    myList.Display()
}
