package main

import (
	"algorithms/data_structures"
	"algorithms/sorting_algorithms"
	"fmt"
)

func main() {
	myArray := []int{4, 2, 5, 6, 23, 54, 34, 23}
	sorting_algorithms.Selection_Sort(&myArray)
	fmt.Println(myArray)
	data_structures.Array_implementation()

	//linked List
	// Creating a linked list
	list := data_structures.LinkedList{}

	// Appending elements to the linked list
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Printing the linked list
	fmt.Println("Linked List:")
	list.PrintList()
}
