package main

import (
	"fmt"
	"github.com/rexposadas/ariel-linked-list/models"
)

func main() {
	// Create a new linked list
	myList := models.LinkedList{}

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
