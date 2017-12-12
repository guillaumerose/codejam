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
	count := 0
	seen := set.New()
	for seed, _ := range table {
		current := group(seed, table)
		if !seen.Has(current.List()[0]) {
			seen.Add(current.List()...)
			count++
		}
	}
	fmt.Println(count)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func group(seed string, table map[string]*set.Set) *set.Set {
	seen := set.New()
	queue := set.New(seed)
	for queue.Size() > 0 {
		current := queue.Pop()
		if seen.Has(current) {
			continue
		}
		seen.Add(current)
		queue.Add(table[current.(string)].List()...)
	}
	return seen
}
