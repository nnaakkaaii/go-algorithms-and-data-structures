# 二分木

与えられた根付き二分木の各節点uについて、次の情報を出力する

- uの節点番号
- uの親
- uの兄弟
- uの子の数
- uの深さ
- uの高さ
- 節点の種類

## 実行例
### 入力例
```
9
0 1 4
1 2 3
2 -1 -1
3 -1 -1
4 5 8
5 6 7
6 -1 -1
7 -1 -1
8 -1 -1
```

### 出力例
```
node 0 : parent = -1. sibling = -1, degree = 2, depth = 0, height = 3, root
...
node 8 : parent = 4, sibling = 5, degree = 0, depth = 2, height = 0, leaf
```

## 考え方
### 二分木の表現
次のような構造体として定義可能

```go
package binary_tree

type Node struct {
	parent int
	left int
	right int
}
```

節点が子を持たない時にはleftとrightはNILとなる

### 節点の高さ

次のような再帰的アルゴリズムにより求めることができる

```go
package binary_tree

import "math"

const NIL = -1

type Node struct {
	parent int
	left   int
	right  int
}

type Tree struct {
	nodes []Node
}

func (t *Tree) getHeight(u int) int {
	h1 := 0
	h2 := 0
	if (*t).nodes[u].right != NIL {
		h1 = (*t).getHeight((*t).nodes[u].right) + 1
	}
	if (*t).nodes[u].left != NIL {
		h2 = (*t).getHeight((*t).nodes[u].left) + 1
	}
	if h1 > h2 {
		return h1
    } else {
    	return h2
    }
}
```

これは`O(n)`の計算量である幅優先探索を行うアルゴリズムとなっている