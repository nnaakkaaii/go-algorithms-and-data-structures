package main

import (
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const k = 100


func countingSort(list []int, numList int) []int {
	counter := make([]int, k)
	sortedList := make([]int, numList)

	// C[Aj]にAjの出現数を記録する
	for j := 0; j < numList; j++ {
		counter[list[j]]++
	}

	// C[i]にi以下の数の出現数を記録する
	for i := 1; i < k; i++ {
		counter[i] = counter[i] + counter[i - 1]
	}

	// B[C[Aj] - 1]にAjの値を格納する
	for j := numList - 1; j >= 0; j-- {
		sortedList[counter[list[j]] - 1] = list[j]
		counter[list[j]]--
	}

	return sortedList
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
	sortedList := countingSort(list, numList)
	fmt.Println(sortedList)
}

