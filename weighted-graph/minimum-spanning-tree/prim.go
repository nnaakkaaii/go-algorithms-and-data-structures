package main

import (
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX = 100
const INFTY = 1<<21
const WHITE = 0
const GRAY = 1
const BLACK = 2

var M = [MAX][MAX]int{}

func prim(n int) int {
	d := make([]int, MAX)  // vを親とする辺で重みが最小の辺の重み
	p := make([]int, MAX)  // MSTにおけるvの親
	color := make([]int, MAX)  // vの訪問状態

	// 初期化
	for i := 0; i < n; i++ {
		d[i] = INFTY
		p[i] = -1
		color[i] = WHITE
	}

	d[0] = 0  // ルートを選択

	for true {
		minv := INFTY
		u := -1
		for i := 0; i < n; i++ {
			// 未探索で最小のd[i]とその時のノードuを記録する
			if minv > d[i] && color[i] != BLACK {
				u = i  // 最初は当然0
				minv = d[i]
			}
		}
		if u == -1 {  // 全てのノードが探索完了
			break
		}
		color[u] = BLACK
		for v := 0; v < n; v++ {
			if color[v] != BLACK && M[u][v] != INFTY {  // 未探索で遷移可能なノード
				if d[v] > M[u][v] {  // uとvの重みが最小の時
					d[v] = M[u][v]  // 最小の重みとして記録
					p[v] = u  // vの親にuを指定
					color[v] = GRAY  // vを未確定のGRAYに
				}
			}
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		if p[i] != -1 {
			sum += M[i][p[i]]
		}
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		list := utils.StrToInt(strings.Split(scanner.Text(), " "))
		for j, elem := range list {
			if elem == -1 {
				M[i][j] = INFTY
			} else {
				M[i][j] = elem
			}
		}
	}

	fmt.Println(prim(n))
}