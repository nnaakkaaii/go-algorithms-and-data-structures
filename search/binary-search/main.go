package main

import (
	"algorithm/utils"
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
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	targetLength, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	target := utils.StrToInt(strings.Split(scanner.Text(), " "))

	scanner.Scan()
	keysLength, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	keys := utils.StrToInt(strings.Split(scanner.Text(), " "))

	for i := 0; i < keysLength; i++ {
		if binarySearch(target, targetLength, keys[i]) >= 0 {
			sum++
		}
	}
	fmt.Printf("%d\n", sum)
}
