package main

import (
	"algorithm/sort/quick-sort/utils/partition"
	"fmt"
)

// binarySearch searches target t from source s, the length of n
func binarySearch(t int, s []int, n int) bool {
	if !sortCheck(s, n) {
		quickSort(&s, n)
	}
	return _binarySearch(t, s, n)
}

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

// sortCheck checks if list lst, the length of n, is sorted
func sortCheck(lst []int, n int) bool {
	for i := 0; i < n-1; i++ {
		if lst[i] > lst[i+1] {
			return false
		}
	}
	return true
}

// quickSort sort list lst, the length of n
func quickSort(lst *[]int, n int) {
	_quickSort(lst, 0, n-1)
}

func _quickSort(lst *[]int, p int, r int) {
	if p < r {
		// リストの最後の要素より小さい部分と大きい部分に分割し、最後の要素のインデックスをqとする
		q := partition.Partition(*lst, p, r)
		// 上で分割したうちの小さい部分を再帰処理
		_quickSort(lst, p, q-1)
		// 上で分割したうちの大きい部分を再帰処理
		_quickSort(lst, q+1, r)
	}
}

func main() {
	a := []int{1, 3, 5, 2, 10, 23, 46, 43, 26, 12, 31, 33, 6, 4, 21, 8, 85, 97}
	n := len(a)
	quickSort(&a, n)
	fmt.Println(a)
	fmt.Println(binarySearch(42, a, n))
}
