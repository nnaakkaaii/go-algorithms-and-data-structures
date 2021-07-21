package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * forの終了条件・キーとの比較の2つの比較演算を行っているため低速
 */
func lineSearchOld(list []string, n int, key string) int {
	for i := 0; i < n; i ++ {
		if list[i] == key {
			return i
		}
	}
	return -1
}

/**
 * 上記アルゴリズムに対して定数倍の高速化が望める
 */
func lineSearch(list []string, n int, key string) int {
	i := 0
	list = append(list, key)
	for list[i] != key {
		i++
	}
	if i == n + 1 {
		return -1
	}
	return i
}


func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	targetLength, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	target := strings.Split(scanner.Text(), " ")
	scanner.Scan()
	keysLength, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	keys := strings.Split(scanner.Text(), " ")

	for i := 0; i < keysLength; i++ {
		if lineSearch(target, targetLength, keys[i]) >= 0 {
			sum++
		}
	}
	fmt.Printf("%d\n", sum)
}
