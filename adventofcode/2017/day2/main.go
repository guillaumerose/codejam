package main

import (
	"strings"
	"strconv"
	"os"
	"bufio"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, a := range strings.Fields(line) {
			i, _ := strconv.Atoi(a)
			divide := divide(line, i)
			if divide != 0 {
				sum += divide
				break
			}
		}
	}
	fmt.Println(sum)
}

func divide(line string, i int) int {
	for _, b := range strings.Fields(line) {
		j, _ := strconv.Atoi(b)
		if i == j {
			continue
		}
		if j%i == 0 {
			return j / i
		}
		if i%j == 0 {
			return i / j
		}
	}
	return 0
}
