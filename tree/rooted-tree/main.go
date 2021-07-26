package main

import (
	tree_components "algorithm/tree/rooted-tree/tree-components"
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

	nodes := make([]tree_components.Node, numNodes)
	for i := 0; i < numNodes; i++ {
		nodes[i] = *tree_components.NewNode()
	}

	for i := 0; i < numNodes; i++ {
		scanner.Scan()
		inputs := utils.StrToInt(strings.Split(scanner.Text(), " "))

		for j := 2; j < 2 + inputs[1]; j++ {
			id := inputs[0]
			if j == 2 {
				// 一番左の子なので、idのleftに格納
				nodes[id].Left = inputs[j]
			}
			nodes[inputs[j]].Parent = id  // 各ノードの親にidを指定
			if j != 1 + inputs[1] {
				// 一番右の子以外は、1つ右をrightに格納
				nodes[inputs[j]].Right = inputs[1 + j]
			}
		}
	}

	tree := tree_components.NewTree(nodes, numNodes)

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
		fmt.Printf("node %d: parent = %d, depth = %d, %s, ", i, tree.GetParent(i), tree.GetDepth(i), s)
		fmt.Println(tree.GetChildren(i))
	}
}
