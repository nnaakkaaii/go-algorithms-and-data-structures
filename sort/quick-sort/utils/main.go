package main

import (
	"algorithm/sort/quick-sort/utils/partition"
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	index := partition.Partition(list, 0, numList - 1)
	fmt.Println(list[:index], list[index], list[1 + index:])
}
