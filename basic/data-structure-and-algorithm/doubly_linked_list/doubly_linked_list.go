package doubly_linked_list

import "fmt"

type DoublyLinkedList struct {
	head  *Node
	tail  *Node
	count int
}
type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (list *DoublyLinkedList) Size() int {
	return list.count
}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.count == 0
}
func (list *DoublyLinkedList) Peek() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	return list.head.value, true
}

func (list *DoublyLinkedList) AddHead(value int) {
	newNode := &Node{value, nil, nil}
	if list.count == 0 {
		list.tail = newNode
		list.head = newNode
	} else {
		list.head.prev = newNode
		newNode.next = list.head
		list.head = newNode
	}
	list.count++
}

func (list *DoublyLinkedList) AddTail(value int) {
	newNode := &Node{value, nil, nil}
	if list.count == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	list.count++
}

func (list *DoublyLinkedList) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	value := list.head.value
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}
	list.count--
	return value, true
}

func (list *DoublyLinkedList) RemoveNode(key int) bool {
	curr := list.head
	if curr == nil { // empty list
		return false
	}
	if curr.value == key { // head is the node with value key.
		curr = curr.next
		list.count--
		if curr != nil {
			list.head = curr
			list.head.prev = nil
		} else {
			list.tail = nil // only one element in list.
		}
		return true
	}
	for curr.next != nil {
		if curr.next.value == key {
			curr.next = curr.next.next
			if curr.next == nil { // last element case.
				list.tail = curr
			} else {
				curr.next.prev = curr
			}
			list.count--
			return true
		}
		curr = curr.next
	}
	return false
}

func (list *DoublyLinkedList) IsPresent(key int) bool {
	temp := list.head
	for temp != nil {
		if temp.value == key {
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *DoublyLinkedList) FreeList() {
	list.tail = nil
	list.head = nil
	list.count = 0
}

func (list *DoublyLinkedList) Print() {
	temp := list.head
	for temp != nil {fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

func (list *DoublyLinkedList) ReverseList() {
	curr := list.head
	var tempNode *Node
	for curr != nil {
		tempNode = curr.next
		curr.next = curr.prev
		curr.prev = tempNode
		if curr.prev == nil {
			list.tail = list.head
			list.head = curr
			return
		}
		curr = curr.prev
	}
	return
}

func (list *DoublyLinkedList) CopyListReversed(dll *DoublyLinkedList) {
	curr := list.head
	for curr != nil {
		dll.AddHead(curr.value)
		curr = curr.next
	}
}