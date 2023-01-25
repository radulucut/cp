// https://codeforces.com/problemset/problem/677/A

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

	s1, _ := reader.ReadString(' ')
	s1 = strings.TrimSpace(s1)
	s2, _ := reader.ReadString('\n')
	s2 = strings.TrimSpace(s2)

	n, _ := strconv.Atoi(s1)
	h, _ := strconv.Atoi(s2)

	l, _ := reader.ReadString('\n')
	l = strings.TrimSpace(l)

	arr := strings.Split(l, " ")
	res := 0

	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(arr[i])

		if a > h {
			res += 2
		} else {
			res++
		}
	}

	fmt.Println(res)
}
