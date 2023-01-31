// https://codeforces.com/problemset/problem/144/A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	arr := strings.Split(s, " ")

	maxH := 0
	minH := 101
	maxPos := 0
	minPos := 0

	for i := 0; i < len(arr); i++ {
		p, _ := strconv.Atoi(arr[i])

		if p > maxH {
			maxH = p
			maxPos = i
		}

		if p <= minH {
			minH = p
			minPos = i
		}
	}

	res := maxPos - 0 + len(arr) - minPos - 1

	if maxPos >= minPos {
		res--
	}

	fmt.Println(res)
}
