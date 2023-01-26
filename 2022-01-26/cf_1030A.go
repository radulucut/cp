// https://codeforces.com/problemset/problem/1030/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')

	s1, _ := reader.ReadString('\n')

	for i := 0; i < len(s1); i++ {
		if s1[i] == '1' {
			fmt.Println("HARD")
			return
		}
	}

	fmt.Println("EASY")
}
