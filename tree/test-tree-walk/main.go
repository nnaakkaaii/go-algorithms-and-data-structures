package main

import (
	binary_tree "algorithm/tree/binary-tree"
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree binary_tree.Tree

func (t *Tree) PreParse(u int) {
	if u == binary_tree.NIL {
		return
	}
	fmt.Printf("%d ", u)
	t.PreParse(t.Nodes[u].Left)
	t.PreParse(t.Nodes[u].Right)
}

func (t *Tree) InParse(u int) {
	if u == binary_tree.NIL {
		return
	}
	t.InParse(t.Nodes[u].Left)
	fmt.Printf("%d ", u)
	t.InParse(t.Nodes[u].Right)
}

func (t *Tree) PostParse(u int) {
	if u == binary_tree.NIL {
		return
	}
	t.PostParse(t.Nodes[u].Left)
	t.PostParse(t.Nodes[u].Right)
	fmt.Printf("%d ", u)
}

func (t *Tree) GetRoot() int {
	for i := 0; i < t.NumNodes; i++ {
		if t.Nodes[i].Parent == binary_tree.NIL {
			return i
		}
	}
	return -1
}

func NewTree(nodes []binary_tree.Node, numNodes int) *Tree {
	tree := new(Tree)
	tree.Nodes = nodes
	tree.NumNodes = numNodes
	tree.Depth = [binary_tree.MAX]int{0}
	return tree
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	numNodes, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("input error")
		return
	}

	// nodesの初期化
	nodes := make([]binary_tree.Node, numNodes)
	for i := 0; i < numNodes; i++ {
		nodes[i] = *binary_tree.NewNode()
	}

	for i := 0; i < numNodes; i++ {
		scanner.Scan()
		inputs := utils.StrToInt(strings.Split(scanner.Text(), " "))
		parent := inputs[0]
		left := inputs[1]
		right := inputs[2]

		nodes[parent].Left = left
		nodes[parent].Right = right
		if left != binary_tree.NIL {
			nodes[left].Parent = parent
		}
		if right != binary_tree.NIL {
			nodes[right].Parent = parent
		}
	}

	tree := NewTree(nodes, numNodes)
	root := tree.GetRoot()
	fmt.Println("\nPreorder")
	tree.PreParse(root)
	fmt.Println("\nInParse")
	tree.InParse(root)
	fmt.Println("\nPostParse")
	tree.PostParse(root)
}

