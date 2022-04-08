package main

import (
	"algorithm/sort/quick-sort/utils/partition"
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var solver = binarySearch

// bruteForce 's calculation order is O(n^4)
func bruteForce(n int, m int, ks []int) bool {
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			for c := 0; c < n; c++ {
				for d := 0; d < n; d++ {
					if ks[a]+ks[b]+ks[c]+ks[d] == m {
						return true
					}
				}
			}
		}
	}
	return false
}

func _quickSort(list *[]int, p int, r int) {
	if p < r {
		// リストの最後の要素より小さい部分と大きい部分に分割し、最後の要素のインデックスをqとする
		q := partition.Partition(*list, p, r)
		// 上で分割したうちの小さい部分を再帰処理
		_quickSort(list, p, q-1)
		// 上で分割したうちの大きい部分を再帰処理
		_quickSort(list, q+1, r)
	}
}

// _binarySearch search target t from source s with the length of n
func _binarySearch(t int, s []int, n int) bool {
	i := n / 2
	if n == 0 {
		return false
	} else if s[i] == t {
		return true
	} else if s[i] < t {
		return _binarySearch(t, s[1+i:], n-1-i)
	} else {
		return _binarySearch(t, s[:i], i)
	}
}

// binarySearch 's calculation order is O(n^2 logn)
func binarySearch(n int, m int, ks []int) bool {
	// ks[c] + ks[d]の取り得る値
	kk := make([]int, n*n)
	for c := 0; c < n; c++ {
		for d := 0; d < n; d++ {
			kk[c*n+d] = ks[c] + ks[d]
		}
	}
	// sort kk
	_quickSort(&kk, 0, n*n-1)
	// binary search : ks[c] + ks[d] = m - ks[a] - ks[b] なる ks[c] + ks[d] がリストに存在するか
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			t := m - ks[a] - ks[b] // t : target of search
			if _binarySearch(t, kk, n*n) {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	r := regexp.MustCompile(`\d+`)

	scanner.Scan()
	ns := r.FindAllString(scanner.Text(), -1)
	if len(ns) != 1 {
		fmt.Println("input error")
		return
	}

	scanner.Scan()
	n, _ := strconv.Atoi(ns[0])
	ms := r.FindAllString(scanner.Text(), -1)
	if len(ms) != 1 {
		fmt.Println("input error")
		return
	}

	scanner.Scan()
	m, _ := strconv.Atoi(ms[0])
	ksText := r.FindAllString(scanner.Text(), -1)
	ks := utils.StrToInt(ksText)

	if solver(n, m, ks) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
