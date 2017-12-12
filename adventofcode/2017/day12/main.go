package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"gopkg.in/fatih/set.v0"
)

func main() {
	table := make(map[string]*set.Set)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		left := strings.Split(text, " <-> ")
		right := strings.Split(left[1], ", ")
		right = append(right, left[0])
		for _, x := range right {
			for _, y := range right {
				if table[x] == nil {
					table[x] = set.New()
				}
				table[x].Add(y)
			}
		}
	}
	seen := set.New()
	queue := set.New("0")
	for queue.Size() > 0 {
		current := queue.Pop()
		if seen.Has(current) {
			continue
		}
		seen.Add(current)
		queue.Add(table[current.(string)].List()...)
	}
	fmt.Println(seen.Size())
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
