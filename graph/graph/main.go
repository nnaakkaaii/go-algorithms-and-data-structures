package main

import (
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100

func main() {
	var M = [N][N]int{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		inputs := utils.StrToInt(strings.Split(scanner.Text(), " "))
		u := inputs[0]  // id
		k := inputs[1]  // 回数
		u--  // 0オリジンへ
		for j := 0; j < k; j++ {
			v := inputs[2 + j]  // 隣接id
			v--  // 0オリジンへ
			M[u][v] = 1
		}
	}

	// 出力
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(M[i][j])
		}
		fmt.Println()
	}
}
