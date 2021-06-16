package doublyLinkedList

import "fmt"

type List struct {
	head  *Node
	count int
}

type Node struct {
	value int
	next  *Node
}

func (list *List) Size() int {
	return list.count
}

func (list *List) IsEmpty() bool {
	return list.count == 0
}

func (list *List) AddHead(value int) {
	list.head = &Node{value, list.head}
	list.count++
}

func (list *List) addTail(value int) {
	curr := list.head
	newNode := &Node{value, nil}
	if curr == nil {
		list.head = newNode
		return
	}
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

func (list *List) SortedInsert(value int) {
	newNode := &Node{value, nil}
	curr := list.head
	if curr == nil || curr.value > value {
		newNode.next = list.head
		list.head = newNode
		return
	}
	for curr.next != nil && curr.next.value < value {
		curr = curr.next
	}
	newNode.next = curr.next
	curr.next = newNode
}

func (list *List) IsPresent(data int) bool {
	temp := list.head
	for temp != nil {
		if temp.value == data {
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *List) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	value := list.head.value
	list.head = list.head.next
	list.count--
	return value, true
}

func (list *List) DeleteNode(delValue int) bool {
	temp := list.head
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return false
	}
	if delValue == list.head.value {
		list.head = list.head.next
		list.count--
		return true
	}
	for temp.next != nil {
		if temp.next.value == delValue {
			temp.next = temp.next.next
			list.count--
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *List) DeleteNodes(delValue int) {
	currNode := list.head
	for currNode != nil && currNode.value == delValue {
		list.head = currNode.next
		currNode = list.head
	}
	for currNode != nil {
		nextNode := currNode.next
		if nextNode != nil && nextNode.value == delValue {
			currNode.next = nextNode.next
		} else {
			currNode = nextNode
		}
	}
}

func (list *List) Reverse() {
	curr := list.head
	var prev, next *Node
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	list.head = prev
}

func (list *List) ReverseRecurse() {
	list.head = list.reverseRecurseUtil(list.head, nil)
}
func (list *List) reverseRecurseUtil(currentNode *Node, nextNode *Node) *Node {
	var ret *Node
	if currentNode == nil {
		return nil
	}
	if currentNode.next == nil {
		currentNode.next = nextNode
		return currentNode
	}
	ret = list.reverseRecurseUtil(currentNode.next, currentNode)
	currentNode.next = nextNode
	return ret
}

func (list *List) LoopDetect() bool {
	slowPtr := list.head
	fastPtr := list.head
	for fastPtr.next != nil && fastPtr.next.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			fmt.Println("loop found")
			return true
		}
	}
	fmt.Println("loop not found")
	return false
}

func (list *List) LoopTypeDetect() int {
	slowPtr := list.head
	fastPtr := list.head
	for fastPtr.next != nil && fastPtr.next.next != nil {
		if list.head == fastPtr.next || list.head == fastPtr.next.next {
			fmt.Println("circular list loop found")
			return 2
		}
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			fmt.Println("loop found")
			return 1
		}
	}
	fmt.Println("loop not found")
	return 0
}

func (list *List) RemoveLoop() {
	loopPoint := list.LoopPointDetect()
	if loopPoint == nil {
		return
	}
	firstPtr := list.head
	if loopPoint == list.head {
		for firstPtr.next != list.head {
			firstPtr = firstPtr.next
		}
		firstPtr.next = nil
		return
	}
	secondPtr := loopPoint
	for firstPtr.next != secondPtr.next {
		firstPtr = firstPtr.next
		secondPtr = secondPtr.next
	}
	secondPtr.next = nil
}
func (list *List) LoopPointDetect() *Node {
	slowPtr := list.head
	fastPtr := list.head
	for fastPtr.next != nil && fastPtr.next.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			return slowPtr
		}
	}
	return nil
}

func (list *List) FindIntersection(head *Node, head2 *Node) *Node {
	l1 := 0
	l2 := 0
	tempHead := head
	tempHead2 := head2
	for tempHead != nil {
		l1++
		tempHead = tempHead.next}
	for tempHead2 != nil {
		l2++
		tempHead2 = tempHead2.next
	}
	var diff int
	if l1 < 12 {
		temp := head
		head = head2
		head2 = temp
		diff = l2 - l1
	} else {
		diff = l1 - l2
	}
	for ; diff > 0; diff-- {
		head = head.next
	}
	for head != head2 {
		head = head.next
		head2 = head2.next
	}
	return head
}