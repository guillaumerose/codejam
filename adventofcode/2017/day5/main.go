package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main() {
	var list []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		list = append(list, i)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(process(list))
}

func process(list []int) int {
	current := 0
	counter := 0
	for {
		if current >= len(list) || current < 0 {
			break
		}
		old := current
		offset := list[current]
		current += offset
		if offset >= 3 {
			list[old]--
		} else {
			list[old]++
		}
		counter++
	}
	return counter
}
