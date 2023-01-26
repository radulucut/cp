// https://codeforces.com/problemset/problem/271/A

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

	s1, _ := reader.ReadString('\n')
	s1 = strings.TrimSpace(s1)

	n, _ := strconv.Atoi(s1)
	res := n

	for true {
		var arr [10]bool
		found := true
		res++
		n = res

		for n > 0 {
			d := n % 10

			if arr[d] {
				found = false
				break
			}

			arr[d] = true
			n /= 10
		}

		if found {
			break
		}
	}

	fmt.Println(res)
}
