package main

import (
	"algorithm/data-structures/queue"
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100
const INFTY = -1

var M = [N][N]int{}
var d = [N]int{}  // 距離で訪問状態を管理する

func bfs(n int, s int) {
	q := queue.NewQueue(N)
	q.Enqueue(s)
	for i := 0; i < n; i++ {
		d[i] = INFTY
	}
	d[s] = 0  // 既に到達できているので

	for !q.IsEmpty() {
		_u, _ := q.Dequeue()  // キューから先入した値を取り出す
		u := _u.(int)
		for v := 0; v < n; v++ {
			if M[u][v] == 0 {  // 到達不可能
				continue
			}
			if d[v] != INFTY {  // 既に距離がわかっている
				continue
			}
			d[v] = d[u] + 1  // 次に到達可能
			q.Enqueue(v)  // 回りたいリストに入れておく
		}
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d %d\n", i + 1, d[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		inputs := utils.StrToInt(strings.Split(scanner.Text(), " "))
		u := inputs[0]  // id
		u--
		k := inputs[1]  // 次数
		for j := 0; j < k; j++ {
			v := inputs[2 + j]  // 隣接要素
			v--
			M[u][v] = 1
		}
	}

	// recursive.Dfs(n)
	bfs(n, 0)
}
