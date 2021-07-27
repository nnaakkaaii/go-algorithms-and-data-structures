# 二分木の巡回

与えられた二分木の全ての節点を体系的に訪問するプログラムを作成する

1. 根節点、左部分木、右部分木の順で節点番号を出力する。これを木の先行順巡回(preorder)と呼ぶ
2. 左部分木、根節点、右部分木の順で節点番号を出力する。これを木の中間順巡回(inorder)と呼ぶ
3. 左部分木、右部分木、根節点の順で節点番号を出力する。これを木の後行順巡回(postorder)と呼ぶ

それぞれは再帰処理により計算が可能

```go
package preorder

import (
	"algorithm/tree/binary-tree"
	"fmt"
)

const NIL = -1

type Tree binary_tree.Tree

func (t *Tree) PreParse(u int) {
	if u == NIL {
		return
	}
	fmt.Println(u)
    (*t).PreParse((*t).Nodes[u].Left)
    (*t).PreParse((*t).Nodes[u].Right)
}

func (t *Tree) InParse(u int) {
	if u == NIL {
		return
    }
    (*t).InParse((*t).Nodes[u].Left)
	fmt.Println(u)
    (*t).InParse((*t).Nodes[u].Right)
}

func (t *Tree) PostParse(u int) {
	if u == NIL {
		return
    }
    (*t).PostParse((*t).Nodes[u].Left)
    (*t).PostParse((*t).Nodes[u].Right)
	fmt.Println(u)
}
```