// https://codeforces.com/problemset/problem/617/A

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

	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	n, _ := strconv.Atoi(str)

	rem := 0

	if n%5 > 0 {
		rem++
	}

	fmt.Printf("%d", n/5+rem)

	fmt.Println()
}
