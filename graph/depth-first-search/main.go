package main

import (
	"algorithm/graph/depth-first-search/stack"
	"algorithm/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

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
			// recursive.M[u][v] = 1
			stack.M[u][v] = 1
		}
	}

	// recursive.Dfs(n)
	stack.Dfs(n)
}
