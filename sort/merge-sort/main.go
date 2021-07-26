package main

import (
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const sentinel = 2000000000

func merge(list []int, left int, mid int, right int) []int {
	n1 := mid - left
	n2 := right - mid

	// スライスの left ~ mid と mid ~ right を新たなスライスにコピーし、末尾にsentinelを加える
	leftList := make([]int, n1)
	rightList := make([]int, n2)
	copy(leftList, list[left:left+n1])
	copy(rightList, list[mid:mid+n2])
	leftList = append(leftList, sentinel)
	rightList = append(rightList, sentinel)

	// スライスのleft ~ rightを整列している (O(n1 + n2)のアルゴリズム)
	i := 0
	j := 0
	for k := left; k < right; k++ {
		if leftList[i] <= rightList[j] {
			list[k] = leftList[i]
			i++
		} else {
			list[k] = rightList[j]
			j++
		}
	}
	fmt.Println(list)
	return list
}

func mergeSort(list []int, numList int, left int, right int) []int {
	if left + 1 < right {
		mid := (left + right) / 2
		// 分割を行っていき、次の2行のいずれもこれ以上分割できなくなったらmergeが実行される
		list = mergeSort(list, numList, left, mid)
		list = mergeSort(list, numList, mid, right)
		list = merge(list, left, mid, right)
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
	sortedList := mergeSort(list, numList, 0, numList)
	fmt.Println(sortedList)
}
