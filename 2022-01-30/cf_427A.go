// https://codeforces.com/problemset/problem/427/A

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

	cops := 0
	res := 0

	for i := 0; i < len(arr); i++ {
		p, _ := strconv.Atoi(arr[i])

		if p > 0 {
			cops += p
		} else {
			p = p * -1

			if p > cops {
				res += p - cops
				cops = 0
			} else {
				cops -= p
			}
		}
	}

	fmt.Println(res)
}
