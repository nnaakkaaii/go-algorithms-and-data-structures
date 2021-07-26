package main

import (
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(list []int, numList int, skipInterval int) []int {
	for i := skipInterval; i < numList; i++ {
		v := list[i]
		j := i - skipInterval
		for j >= 0 && list[j] > v {
			list[j+skipInterval] = list[j]
			j = j - skipInterval
		}
		list[j+skipInterval] = v
	}
	return list
}

func shellSort(list []int, numList int) []int {
	var skipIntervalList []int
	skipInterval := 1
	for skipInterval < numList {
		skipIntervalList = append([]int{skipInterval}, skipIntervalList...)
		skipInterval = 3*skipInterval + 1
	}
	for _, skipInterval = range skipIntervalList {
		insertionSort(list, numList, skipInterval)
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
	sortedList := shellSort(list, numList)
	fmt.Println(sortedList)
}
