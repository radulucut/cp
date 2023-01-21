// https://codeforces.com/problemset/problem/977/A

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

	n, _ := reader.ReadString(' ')
	kStr, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	kStr = strings.TrimSpace(kStr)
	k, _ := strconv.Atoi(kStr)

	i := len(n) - 1

	for k > 0 {
		d := int(n[i] - '0')

		if k < d {
			break
		}

		k -= d + 1
		i--
	}

	for j := 0; j < i; j++ {
		fmt.Printf("%c", n[j])
	}

	if k > 0 {
		fmt.Print(string(int(n[i]) - k))
	} else {
		fmt.Print(string(n[i]))

		if k == -1 {
			fmt.Print("0")
		}
	}

	fmt.Println()
}
