// 挿入が可能
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Key int
	Parent *Node
	Left *Node
	Right *Node
}

type Tree struct {
	Root *Node
	NIL *Node
}

func NewTree() *Tree {
	tree := new(Tree)
	Nil := new(Node)
	(*tree).NIL = Nil
	(*tree).Root = Nil
	return tree
}

func (t *Tree) Insert(key int) {
	parent := (*t).NIL
	leaf := (*t).Root

	z := new(Node)
	(*z).Key = key
	(*z).Left = (*t).NIL
	(*z).Right = (*t).NIL

	for leaf != (*t).NIL {
		parent = leaf
		if (*z).Key < (*leaf).Key {
			leaf = (*leaf).Left
		} else {
			leaf = (*leaf).Right
		}
	}

	(*z).Parent = parent

	if parent == (*t).NIL {
		(*t).Root = z
	} else {
		if (*z).Key < (*parent).Key {
			(*parent).Left = z
		} else {
			(*parent).Right = z
		}
	}
}

func (t *Tree) InOrder(u *Node) {
	if u == (*t).NIL {
		return
	}
	(*t).InOrder((*u).Left)
	fmt.Printf("%d ", (*u).Key)
	(*t).InOrder((*u).Right)
}

func (t *Tree) PreOrder(u *Node) {
	if u == (*t).NIL {
		return
	}
	fmt.Printf("%d ", (*u).Key)
	(*t).PreOrder((*u).Left)
	(*t).PreOrder((*u).Right)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numCommands, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("input error")
		return
	}

	tree := NewTree()

	for i := 0; i < numCommands; i++ {
		scanner.Scan()
		command := strings.Split(scanner.Text(), " ")

		if command[0] == "insert" {
			key, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Println("input error")
				return
			}
			tree.Insert(key)
		} else if command[0] == "print" {
			fmt.Println("\nInOrder")
			tree.InOrder(tree.Root)
			fmt.Println("\nPreOrder")
			tree.PreOrder(tree.Root)
		} else {
			fmt.Println("input error")
			return
		}
	}
}
