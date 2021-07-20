package main

import (
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
	var intList []int
	for _, strElem := range strings.Split(scanner.Text(), " ") {
		intElem, err := strconv.Atoi(strElem)
		if err != nil {
			fmt.Println("input error")
			return
		}
		intList = append(intList, intElem)
	}
	sortedList := insertionSort(intList, numList)
	fmt.Println(sortedList)
}
