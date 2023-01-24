// https://codeforces.com/problemset/problem/41/A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s1, _ := reader.ReadString('\n')
	s1 = strings.TrimSpace(s1)
	s2, _ := reader.ReadString('\n')
	s2 = strings.TrimSpace(s2)

	ln := len(s1)

	if ln != len(s2) {
		fmt.Println("NO")
		return
	}

	for i := 0; i < ln; i++ {
		if s1[i] != s2[ln-i-1] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
