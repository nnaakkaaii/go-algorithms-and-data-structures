package partition

// Partition は list の要素をrより小さい要素, r, rより大きい要素に並び替え、rのインデックスを返す
func Partition(list []int, p int, r int) int {
	x := list[r]
	i := p - 1
	for j := p; j < r; j++ {
		if list[j] <= x {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i + 1], list[r] = list[r], list[i + 1]
	return i + 1
}
