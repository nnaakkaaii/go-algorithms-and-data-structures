package main

import (
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func selectionSort(list []int, numList int) []int {
	for i := 0; i < numList; i++ {
		minJ := i
		for j := i; j < numList; j++ {
			if list[j] < list[minJ] {
				minJ = j
			}
		}
		list[i], list[minJ] = list[minJ], list[i]
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
	sortedList := selectionSort(list, numList)
	fmt.Println(sortedList)
}
