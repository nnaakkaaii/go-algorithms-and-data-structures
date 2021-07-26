# 根付き木の表現

与えられた根付き木Tの各node uについて、次の情報を出力するプログラムを作成する

- uの節点番号
- uの節点の種類
- uの親の節点番号
- uの子のリスト
- uの深さ

## 実行例
### 入力例

1つ目の要素がid, 2つ目の要素が次数, それ以降が連結するノードのid

```
13
0 3 1 4 10
1 2 2 3
2 0
3 0
4 3 5 6 7
5 0
6 0
7 2 8 9
8 0
9 0
10 2 11 12
11 0
12 0
```

### 出力例
```
node 0 : parent = -1, depth = 0, root, [1, 4, 10]
node 1 : parent = 0, depth = 1, internal node, [2, 3]
node 2 : parent = 1, depth = 2, leaf, []
...
node 11 : parent = 10, depth = 2, leaf, []
node 12 : parent = 10, depth = 2, leaf, []
```

## 考え方
### ノードの実装

入力が完了した後はnodeの数が変化しないので、各nodeが次の情報を持つような「左子右兄弟表現」により木を表すことにする。

- ノードuの親
- ノードuの最も左の子
- ノードuのすぐ右の兄弟

```go
package tree

type Node struct {
	parent int
	left int
	right int
}
```

この時、
- ノードuの親は`u.parent`によりわかり、存在しない場合はroot
- `u.left`が存在しないnodeがleaf
- `u.right`が存在しないnodeが最も右の子

また、親、左の子、右の兄弟が存在しないことを示すためにデフォルト値として特別な値NILを用意する。

### 節点の深さ
節点の深さは、uから親を辿っていき根に至るまでの辺の数を数える。
根の親をNIL(=-1)にしておくことで、他の節点との区別が可能。

```go
package tree

const NIL = -1

type Node struct {
	parent int
	left int
	right int
}

type T []Node

func (t *T) getDepth(u int) int {
	d := 0
	for (*t)[u].parent != NIL {
		u = (*t)[u].parent
		d++
    }
    return d
}
```

一方で、次のように再帰的なアルゴリズムで木の節点の深さを高速に求めることが可能

```go
package tree

const MAX = 100005
const NIL = -1
var D = [MAX]int{-1}

type Node struct {
	parent int
	left int
	right int
}

type T []Node

func (t *T) setDepth(u int, p int) {
	D[u] = p  // ノードuの深さを格納する配列に代入
	if (*t)[u].right != NIL {
        (*t).setDepth((*t)[u].right, p)
    }
    if (*t)[u].left != NIL {
        (*t).setDepth((*t)[u].left, p + 1)
    }
}
```

このアルゴリズムでは、右の兄弟の深さと最も左の子の深さを再帰的に計算している。

前者のアルゴリズムは`O(nh)`なのに対し、後者は`O(n)`のアルゴリズムとなる