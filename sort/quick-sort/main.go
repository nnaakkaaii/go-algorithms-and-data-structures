package main

import (
	"algorithm/sort/quick-sort/utils/partition"
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func quickSort(list []int, p int, r int) []int {
	if p < r {
		// リストの最後の要素より小さい部分と大きい部分に分割し、最後の要素のインデックスをqとする
		q := partition.Partition(list, p, r)
		// 上で分割したうちの小さい部分を再帰処理
		list = quickSort(list, p, q - 1)
		// 上で分割したうちの大きい部分を再帰処理
		list = quickSort(list, q + 1, r)
	}
	return list
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// リストの個数を入力
	scanner.Scan()
	numList, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("input error")
		return
	}
	// リストを入力
	scanner.Scan()
	list := utils.StrToInt(strings.Split(scanner.Text(), " "))
	list = quickSort(list, 0, numList - 1)
	fmt.Println(list)
}

