// https://codeforces.com/contest/59/problem/A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	uppercaseCount := 0
	for i := 0; i < len(str); i++ {
		if str[i] < 97 {
			uppercaseCount++
		}
	}

	doUppercase := uppercaseCount > len(str)-uppercaseCount

	for i := 0; i < len(str); i++ {
		if doUppercase && str[i] >= 97 {
			fmt.Print(string(str[i] - 32))

			continue
		}

		if !doUppercase && str[i] < 97 {
			fmt.Print(string(str[i] + 32))

			continue
		}

		fmt.Printf("%c", str[i])
	}

	fmt.Println()
}
