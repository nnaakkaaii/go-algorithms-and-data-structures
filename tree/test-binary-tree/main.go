package main

import (
	"algorithm/tree/binary-tree"
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NIL = -1

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
		if left != NIL {
			nodes[left].Parent = parent
		}
		if right != NIL {
			nodes[right].Parent = parent
		}
	}

	tree := binary_tree.NewTree(nodes, numNodes)

	for i := 0; i < numNodes; i++ {
		s := ""
		switch {
		case tree.IsRoot(i):
			s = "root"
			break
		case tree.IsLeaf(i):
			s = "leaf"
			break
		case tree.IsInternalNode(i):
			s = "internal node"
			break
		}
		fmt.Printf(
			"node %d: parent = %d, sibling = %d, depth = %d, height = %d, %s\n",
			i, tree.GetParent(i), tree.GetSibling(i), tree.GetDepth(i), tree.GetHeight(i), s,
		)
	}
}
