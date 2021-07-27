package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var dp = [50]int{0}

func initializeDp(initVal int) {
	for i, _ := range dp {
		dp[i] = initVal
	}
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		dp[n] = 1
		return dp[n]
	}
	if dp[n] != -1 {
		return dp[n]
	}
	dp[n] = fibonacci(n - 1) + fibonacci(n - 2)
	return dp[n]
}

func main() {
	initializeDp(-1)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	fmt.Println(fibonacci(num))
}
