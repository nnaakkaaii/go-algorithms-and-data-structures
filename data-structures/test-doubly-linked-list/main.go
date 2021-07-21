package main

import doubly_linked_list "algorithm/data-structures/doubly-linked-list"


func main() {
	doubleLinkedList := doubly_linked_list.NewDoubleLinkedList()
	doubleLinkedList.Insert(1)
	doubleLinkedList.Insert(2)
	doubleLinkedList.Insert(3)
	doubleLinkedList.PrintList()  // 3 2 1
	doubleLinkedList.DeleteFirst()
	doubleLinkedList.PrintList()  // 2 1
	doubleLinkedList.Insert(4)
	doubleLinkedList.DeleteLast()
	doubleLinkedList.PrintList()  // 4 2
	doubleLinkedList.Insert(5)
	doubleLinkedList.PrintList()  // 5 4 2
	doubleLinkedList.DeleteKey(2)
	doubleLinkedList.PrintList()  // 5 4
}
