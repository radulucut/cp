// https://codeforces.com/problemset/problem/734/A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			count++
		}
	}

	rem := len(s) - count
	if count > rem {
		fmt.Println("Anton")
	} else if count == rem {
		fmt.Println("Friendship")
	} else {
		fmt.Println("Danik")
	}
}
