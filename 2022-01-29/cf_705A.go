// https://codeforces.com/problemset/problem/705/A

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

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print("I hate ")
		} else {
			fmt.Print("I love ")
		}

		if i != n {
			fmt.Print("that ")
		}
	}

	fmt.Print("it\n")
}
