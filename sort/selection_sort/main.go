package main

import (
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
	var intList []int
	for _, strElem := range strings.Split(scanner.Text(), " ") {
		intElem, err := strconv.Atoi(strElem)
		if err != nil {
			fmt.Println("input error")
			return
		}
		intList = append(intList, intElem)
	}
	sortedList := selectionSort(intList, numList)
	fmt.Println(sortedList)
}
