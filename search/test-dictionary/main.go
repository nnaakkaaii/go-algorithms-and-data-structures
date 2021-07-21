package main

import (
	"algorithm/search/dictionary"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dict := dictionary.NewDict()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		operation, keyString := input[0], input[1]

		if operation == "insert" {
			dict.Insert(keyString)
		} else {
			if dict.Find(keyString) {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}
	}
}
