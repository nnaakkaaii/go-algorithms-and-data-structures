package main

import (
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n int
var list []int

func solve (i int, m int) bool {
	if m == 0 {
		return true
	} else if i >= n {
		return false
	}
	return solve(i + 1, m) || solve(i + 1, m - list[i])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	list = utils.StrToInt(strings.Split(scanner.Text(), " "))

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	keys := utils.StrToInt(strings.Split(scanner.Text(), " "))

	for i := 0; i < q; i++ {
		if solve(0, keys[i]) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
