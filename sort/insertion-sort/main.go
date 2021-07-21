package main

import (
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(list []int, numList int) []int {
	for i := 1; i < numList; i++ {
		v := list[i]
		j := i - 1
		for j >= 0 && list[j] > v {
			list[j + 1] = list[j]
			j--
		}
		list[j + 1] = v
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
	sortedList := insertionSort(list, numList)
	fmt.Println(sortedList)
}
