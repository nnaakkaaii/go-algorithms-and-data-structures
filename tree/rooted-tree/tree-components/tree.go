package tree_components

const MAX = 100005
const NIL = -1

type Node struct {
	Parent int
	Left int
	Right int
}

func NewNode() *Node {
	node := new(Node)
	node.Parent = NIL
	node.Left = NIL
	node.Right = NIL
	return node
}

type Tree struct {
	nodes []Node
	numNodes int
	depth [MAX]int
}

func NewTree(nodes []Node, numNodes int) *Tree {
	tree := new(Tree)
	tree.nodes = nodes
	tree.numNodes = numNodes
	tree.depth = [MAX]int{0}
	r := tree.getRoot()
	tree.setDepth(r, 0)  // depthに深さを代入しておく
	return tree
}

func (t *Tree) getRoot() int {
	for i := 0; i < t.numNodes; i++ {
		if t.nodes[i].Parent == NIL {
			return i
		}
	}
	return -1
}

func (t *Tree) setDepth(u int, p int) {
	t.depth[u] = p
	if t.nodes[u].Right != NIL {
		t.setDepth(t.nodes[u].Right, p)
	}
	if t.nodes[u].Left != NIL {
		t.setDepth(t.nodes[u].Left, p + 1)
	}
}

func (t *Tree) GetDepth(u int) int {
	return t.depth[u]
}

func (t *Tree) GetParent(u int) int {
	return t.nodes[u].Parent
}

func (t *Tree) GetChildren(u int) []int {
	c := t.nodes[u].Left
	var ret []int
	for c != NIL {
		ret = append(ret, c)
		c = t.nodes[c].Right
	}
	return ret
}

func (t *Tree) IsRoot(u int) bool {
	return t.GetParent(u) == NIL
}

func (t *Tree) IsLeaf(u int) bool {
	return t.nodes[u].Left == NIL
}

func (t *Tree) IsInternalNode(u int) bool {
	return !t.IsRoot(u) && !t.IsLeaf(u)  // 葉でも根でもない
}

