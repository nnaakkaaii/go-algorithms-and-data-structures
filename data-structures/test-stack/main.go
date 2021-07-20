package main

import (
	"algorithm/data-structures/stack"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := stack.NewStack(200)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var a int
	var b int
	var intElem int
	var err error

	for _, strElem := range strings.Split(scanner.Text(), " ") {
		if strElem == "+" {
			a, err = s.Pop()
			if err != nil {
				goto Error
			}
			b, err = s.Pop()
			err := s.Push(a + b)
			if err != nil {
				goto Error
			}
		} else if strElem == "-" {
			a, err = s.Pop()
			if err != nil {
				goto Error
			}
			b, err = s.Pop()
			if err != nil {
				goto Error
			}
			err = s.Push(b - a)
			if err != nil {
				goto Error
			}
		} else if strElem == "*" {
			a, err = s.Pop()
			if err != nil {
				goto Error
			}
			b, err = s.Pop()
			if err != nil {
				goto Error
			}
			err = s.Push(a * b)
			if err != nil {
				goto Error
			}
		} else {
			intElem, err = strconv.Atoi(strElem)
			if err != nil {
				goto Error
			}
			err = s.Push(intElem)
			if err != nil {
				goto Error
			}
		}
	}

	a, err = s.Pop()
	if err != nil {
		goto Error
	}
	fmt.Printf("= %d\n", a)
	return

Error:
	fmt.Println("input error")
	return
}
