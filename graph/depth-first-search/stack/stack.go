package stack

import (
	stack2 "algorithm/data-structures/stack"
	"fmt"
)

const N = 100
const WHITE = 0  // 未訪問
const GRAY = 1  // 隣接頂点を訪問中
const BLACK = 2  // 訪問済

var n = 0  // ノード数, 後から変更
var M = [N][N]int{}
var color = [N]int{}
var d = [N]int{} // vを最初に訪問した発見時刻
var f = [N]int{} // vの隣接リストを調べ終えた完了時刻
var nt = [N]int{}  // ?
var tt = 0       // 時刻

// uに隣接するvを番号順に取得
func next(u int) int {
	for v := nt[u]; v < n; v++ {
		nt[u] = v + 1
		if M[u][v] > 0 {
			return v
		}
	}
	return -1
}

// 再帰関数による深さ優先探索
func dfsVisit(u int) {
	nt = [N]int{0}
	stack := stack2.NewStack(N)
	stack.Push(u)
	color[u] = GRAY
	tt++
	d[u] = tt

	for !stack.IsEmpty() {
		u := stack.Top().(int)
		v := next(u)
		if v != -1 {
			if color[v] == WHITE {
				color[v] = GRAY
				tt++
				d[v] = tt
				stack.Push(v)
			}
		} else {
			stack.Pop()
			color[u] = BLACK
			tt++
			f[u] = tt
		}
	}
}

func Dfs(numNodes int) {
	n = numNodes
	for u := 0; u < n; u++ {
		// 未訪問のuを始点として深さ優先探索
		if color[u] == WHITE {
			dfsVisit(u)
		}
	}
	for u := 0; u < n; u++ {
		fmt.Printf("%d %d %d\n", u + 1, d[u], f[u])
	}
}
