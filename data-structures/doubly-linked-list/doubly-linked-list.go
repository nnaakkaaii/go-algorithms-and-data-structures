package doubly_linked_list

import "fmt"

type Node struct {
	key int
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	sentinel *Node
}

func NewDoubleLinkedList() *DoubleLinkedList {
	doubleLinkedList := new(DoubleLinkedList)
	sentinel := new(Node)
	(*sentinel).prev = sentinel
	(*sentinel).next = sentinel
	(*doubleLinkedList).sentinel = sentinel
	return doubleLinkedList
}

func (doubleLinkedList *DoubleLinkedList) listSearch(key int) *Node {
	cur := (*(*doubleLinkedList).sentinel).next
	for cur != (*doubleLinkedList).sentinel && (*cur).key != key {
		cur = (*cur).next
	}
	return cur
}

func (doubleLinkedList *DoubleLinkedList) deleteNode(node *Node) {
	if node == (*doubleLinkedList).sentinel {
		return
	}
	(*(*node).prev).next = (*node).next
	(*(*node).next).prev = (*node).prev
	node = nil
}

func (doubleLinkedList *DoubleLinkedList) DeleteFirst() {
	doubleLinkedList.deleteNode((*(*doubleLinkedList).sentinel).next)
}

func (doubleLinkedList *DoubleLinkedList) DeleteLast() {
	doubleLinkedList.deleteNode((*(*doubleLinkedList).sentinel).prev)
}

func (doubleLinkedList *DoubleLinkedList) DeleteKey(key int) {
	doubleLinkedList.deleteNode(doubleLinkedList.listSearch(key))
}

func (doubleLinkedList *DoubleLinkedList) Insert(key int) {
	node := new(Node)
	(*node).key = key
	(*node).next = (*(*doubleLinkedList).sentinel).next
	(*(*(*doubleLinkedList).sentinel).next).prev = node
	(*(*doubleLinkedList).sentinel).next = node
	(*node).prev = (*doubleLinkedList).sentinel
}

func (doubleLinkedList *DoubleLinkedList) PrintList() {
	cur := (*(*doubleLinkedList).sentinel).next
	isf := 0
	for true {
		if cur == (*doubleLinkedList).sentinel {
			break
		}
		isf++
		if isf > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", (*cur).key)
		cur = (*cur).next
	}
	fmt.Println()
}
