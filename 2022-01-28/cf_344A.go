// https://codeforces.com/problemset/problem/344/A

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

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	n, _ := strconv.Atoi(s)

	last, _ := reader.ReadString('\n')
	last = strings.TrimSpace(last)

	res := 1

	for i := 1; i < n; i++ {
		p, _ := reader.ReadString('\n')
		p = strings.TrimSpace(p)

		if p[0] == last[1] {
			last = p
			res++
		}
	}

	fmt.Println(res)
}
