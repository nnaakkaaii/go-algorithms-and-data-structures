package main

import (
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSortOld(list []int, numList int) []int {
	flag := true
	for flag {
		flag = false
		for j := numList - 1; j > 0; j-- {
			if list[j] < list[j - 1] {
				list[j], list[j - 1] = list[j - 1], list[j]
				flag = true
			}
		}
	}
	return list
}

func bubbleSort(list []int, numList int) []int {
	flag := true
	i := 0  // 未ソート部分列のインデックス
	for flag {
		flag = false
		for j := numList - 1; j > i; j-- {
			if list[j] < list[j - 1] {
				list[j], list[j - 1] = list[j - 1], list[j]
				flag = true
			}
		}
		i++
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
	sortedList := bubbleSort(list, numList)
	fmt.Println(sortedList)
}
