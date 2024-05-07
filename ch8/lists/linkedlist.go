package lists

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
}

func (ll *LinkedList[T]) Add(value T) {
	node := Node[T]{Value: value}
	if ll.Head == nil {
		ll.Head = &node
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &node
}

func (ll *LinkedList[T]) Insert(value T, index int) {
	currentPos := 0
	currentNode := ll.Head
	for currentPos < index-1 {
		currentNode = currentNode.Next
		currentPos++
	}
	newNode := Node[T]{Value: value}
	newNode.Next = currentNode.Next
	currentNode.Next = &newNode
}

func (ll *LinkedList[T]) Index(value T) int {
	currentNode := ll.Head
	currentPos := 0
	for currentNode.Next != nil {
		if currentNode.Value == value {
			return currentPos
		}
		currentNode = currentNode.Next
		currentPos++
	}
	return -1
}
func TestLinkedList() {
	ll := LinkedList[int]{}

	// Add elements to the linked list
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	ll.Add(6)
	ll.Add(7)
	ll.Add(8)
	ll.Add(9)
	ll.Add(10)

	// Insert an element at index 5
	ll.Insert(100, 5)

	// Get the indexes of two elements
	index1 := ll.Index(3)
	index2 := ll.Index(100)

	fmt.Println("Index of 3:", index1)
	fmt.Println("Index of 100:", index2)
}
