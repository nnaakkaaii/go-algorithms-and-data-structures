package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binarySearch(list []int, listLength int, key int) int {
	left := 0
	right := listLength
	for left < right {
		mid := (left + right) / 2
		if list[mid] == key {
			return mid
		} else if key < list[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	strToInt := func (strList []string) []int {
		intList := make([]int, len(strList))
		for i, strElem := range strList {
			intElem, _ := strconv.Atoi(strElem)
			intList[i] = intElem
		}
		return intList
	}
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	targetLength, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	targetStr := strings.Split(scanner.Text(), " ")
	targetInt := strToInt(targetStr)

	scanner.Scan()
	keysLength, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	keysStr := strings.Split(scanner.Text(), " ")
	keysInt := strToInt(keysStr)

	for i := 0; i < keysLength; i++ {
		if binarySearch(targetInt, targetLength, keysInt[i]) >= 0 {
			sum++
		}
	}
	fmt.Printf("%d\n", sum)
}
