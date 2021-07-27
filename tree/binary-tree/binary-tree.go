package binary_tree

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
	Nodes []Node
	NumNodes int
	Depth [MAX]int
}

func NewTree(nodes []Node, numNodes int) *Tree {
	tree := new(Tree)
	tree.Nodes = nodes
	tree.NumNodes = numNodes
	tree.Depth = [MAX]int{0}
	r := tree.GetRoot()
	tree.SetDepth(r, 0)  // Depthに深さを代入しておく
	return tree
}

func (t *Tree) GetRoot() int {
	for i := 0; i < t.NumNodes; i++ {
		if t.Nodes[i].Parent == NIL {
			return i
		}
	}
	return -1
}

func (t *Tree) SetDepth(u int, p int) {
	t.Depth[u] = p
	if t.Nodes[u].Right != NIL {
		t.SetDepth(t.Nodes[u].Right, p)
	}
	if t.Nodes[u].Left != NIL {
		t.SetDepth(t.Nodes[u].Left, p + 1)
	}
}

func (t *Tree) GetDepth(u int) int {
	return t.Depth[u]
}

// GetHeight を追加
func (t *Tree) GetHeight(u int) int {
	h1 := 0
	h2 := 0
	if (*t).Nodes[u].Right != NIL {
		h1 = (*t).GetHeight((*t).Nodes[u].Right) + 1
	}
	if (*t).Nodes[u].Left != NIL {
		h2 = (*t).GetHeight((*t).Nodes[u].Left) + 1
	}
	if h1 > h2 {
		return h1
	} else {
		return h2
	}
}

func (t *Tree) GetParent(u int) int {
	return t.Nodes[u].Parent
}

// GetSibling を追加
func (t *Tree) GetSibling(u int) int {
	parent := t.Nodes[u].Parent
	if parent == NIL {
		return NIL
	}
	if t.Nodes[parent].Left != u && t.Nodes[parent].Left != NIL {
		// 親ノードの左に自分やNIL以外の要素がある場合
		return t.Nodes[parent].Left
	}
	if t.Nodes[parent].Right != u && t.Nodes[parent].Right != NIL {
		return t.Nodes[parent].Right
	}
	return NIL
}

func (t *Tree) GetChildren(u int) []int {
	c := t.Nodes[u].Left
	var ret []int
	for c != NIL {
		ret = append(ret, c)
		c = t.Nodes[c].Right
	}
	return ret
}

func (t *Tree) IsRoot(u int) bool {
	return t.GetParent(u) == NIL
}

func (t *Tree) IsLeaf(u int) bool {
	return t.Nodes[u].Left == NIL
}

func (t *Tree) IsInternalNode(u int) bool {
	return !t.IsRoot(u) && !t.IsLeaf(u)  // 葉でも根でもない
}

