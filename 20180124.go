package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	args := nextLine()
	cols := strings.Split(args, " ")

	for i := len(cols) - 1; i >= 0; i-- {
		fmt.Println(cols[i])
	}
}
