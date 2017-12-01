package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main() {
	fmt.Println(process(os.Args[1]))
}

func process(input string) int {
	split := strings.Split(input, "")
	acc := 0
	for i := 0; i < len(split); i++ {
		current := split[i]
		next := split[(i+len(split)/2)%len(split)]
		if current == next {
			val, _ := strconv.Atoi(current)
			acc += val
		}
	}
	return acc
}
