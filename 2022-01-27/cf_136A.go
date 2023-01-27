// https://codeforces.com/problemset/problem/136/A

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
	var res [101]int

	for i := 0; i < len(arr); i++ {
		p, _ := strconv.Atoi(arr[i])
		res[p] = i + 1
	}

	for i := 1; i <= len(arr); i++ {
		fmt.Printf("%d ", res[i])
	}

	fmt.Println()
}
