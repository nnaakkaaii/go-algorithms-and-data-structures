package main

import (
	"algorithm/data-structures/queue"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Process struct {
	name string
	time int
}

func min(a int, b int) int {
	return map[bool]int{true: a, false: b}[a < b]
}

func main() {
	q := queue.NewQueue(1000)
	elaps := 0

	// プロセスをキューに順番に追加する
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	nProcess, _ := strconv.Atoi(input[0])
	useTime, _ := strconv.Atoi(input[1])

	for i := 0; i < nProcess; i++ {
		scanner.Scan()
		input = strings.Split(scanner.Text(), " ")
		process := new(Process)
		process.name = input[0]
		process.time, _ = strconv.Atoi(input[1])
		_ = q.Enqueue(*process)
	}

	for !q.IsEmpty() {
		any, _ := q.Dequeue()
		u, _ := any.(Process)
		c := min(useTime, u.time)
		u.time -= c
		elaps += c
		if u.time > 0 {
			_ = q.Enqueue(u)
		} else {
			fmt.Printf("%s %d\n", u.name, elaps)
		}
	}
}
