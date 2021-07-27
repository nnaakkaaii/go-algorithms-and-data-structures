package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 1000

func LCS(X string, Y string) int {
	var c = [N + 1][N + 1]int{}

	m := len(X)
	n := len(Y)

	maxl := 0
	X = " " + X
	Y = " " + Y

	max := func (x int, y int) int {return map[bool]int{true: x, false: y}[x > y]}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i] == Y[j] {
				c[i][j] = c[i - 1][j - 1] + 1
			} else {
				c[i][j] = max(c[i - 1][j], c[i][j - 1])
			}
			maxl = max(maxl, c[i][j])
		}
	}
	return maxl
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < num * 2; i++ {
		scanner.Scan()
		stringA := scanner.Text()
		scanner.Scan()
		stringB := scanner.Text()

		lcs := LCS(stringA, stringB)

		fmt.Println(lcs)
	}
}