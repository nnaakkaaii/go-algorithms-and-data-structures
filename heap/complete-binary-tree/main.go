package main

import (
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2 * i + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	numList, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("input error")
		return
	}

	scanner.Scan()
	list := utils.StrToInt(strings.Split(scanner.Text(), " "))
	for i := 1; i <= numList; i++ {
		fmt.Printf("node %d: ", i)
		fmt.Printf("key = %d, ", list[i - 1])
		if parent(i) >= 1 {
			fmt.Printf("parent key = %d, ", list[parent(i) - 1])
		}
		if left(i) <= numList {
			fmt.Printf("left key = %d, ", list[left(i) - 1])
		}
		if right(i) <= numList {
			fmt.Printf("right key = %d, ", list[right(i) - 1])
		}
		fmt.Println()
	}
}
