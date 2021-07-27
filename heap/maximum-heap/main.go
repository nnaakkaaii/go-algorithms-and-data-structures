package main

import (
	"algorithm/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Heap struct {
	list []int
	numList int
}

func NewHeap(list []int, numList int) *Heap {
	heap := new(Heap)
	heap.list = list
	heap.numList = numList
	return heap
}

func (h *Heap) left(i int) int {
	return 2 * i
}

func (h *Heap) right(i int) int {
	return 2 * i + 1
}

func (h *Heap) maxHeapify(i int) {
	left := (*h).left(i)
	right := (*h).right(i)

	// 左の子、自分、右の子で値が最大のノードを選ぶ
	largest := 0
	if left <= (*h).numList && (*h).list[left - 1] > (*h).list[i - 1] {
		largest = left
	} else {
		largest = i
	}
	if right <= (*h).numList && (*h).list[right - 1] > (*h).list[largest - 1] {
		largest = right
	}
	if largest != i {
		(*h).list[i - 1], (*h).list[largest - 1] = (*h).list[largest - 1], (*h).list[i - 1]
		(*h).maxHeapify(largest)
	}
}

func (h *Heap) buildMaxHeap() {
	for i := (*h).numList / 2; i >= 1; i-- {
		(*h).maxHeapify(i)
	}
}

func (h *Heap) print() {
	fmt.Println((*h).list)
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

	heap := NewHeap(list, numList)
	heap.buildMaxHeap()
	heap.print()
}
