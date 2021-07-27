package recursive

import (
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
var tt = 0       // 時刻

// 再帰関数による深さ優先探索
func dfsVisit(u int) {
	color[u] = GRAY
	tt++
	d[u] = tt
	for v := 0; v < n; v++ {
		if M[u][v] == 0 { // uから到達不可能
			continue
		}
		if color[v] == WHITE { // 未訪問
			dfsVisit(v) // 訪問
		}
	}
	color[u] = BLACK // 訪問済
	tt++
	f[u] = tt
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
