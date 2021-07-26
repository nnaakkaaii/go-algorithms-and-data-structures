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

func merge(list []int, left int, mid int, right int) int {
	n1 := mid - left
	n2 := right - mid

	cnt := 0  // counter

	// スライスの left ~ mid と mid ~ right を新たなスライスにコピーし、末尾にsentinelを加える
	leftList := make([]int, n1)
	rightList := make([]int, n2)
	copy(leftList, list[left:left+n1])
	copy(rightList, list[mid:mid+n2])
	leftList = append(leftList, sentinel)
	rightList = append(rightList, sentinel)

	i := 0
	j := 0
	for k := left; k < right; k++ {
		if leftList[i] <= rightList[j] {
			list[k] = leftList[i]
			i++
		} else {
			list[k] = rightList[j]
			j++
			cnt += n1 - i  // 何個飛ばしたか
		}
	}
	return cnt
}

func mergeSort(list []int, left int, right int) int {
	if left + 1 < right {
		mid := (left + right) / 2
		v1 := mergeSort(list, left, mid)
		v2 := mergeSort(list, mid, right)
		v3 := merge(list, left, mid, right)
		return v1 + v2 + v3
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	numList, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("input error")
		return
	}

	// リストを入力
	scanner.Scan()
	list := utils.StrToInt(strings.Split(scanner.Text(), " "))

	cnt := mergeSort(list, 0, numList)
	fmt.Println(cnt)
}