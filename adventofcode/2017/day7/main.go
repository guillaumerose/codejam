package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	tree := make(map[string]string)
	entrypoint := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		if len(text) > 3 {
			for i := 3; i < len(text); i++ {
				replace := strings.Trim(text[i], ",")
				tree[replace] = text[0]
				entrypoint = replace
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(parent(tree, entrypoint))
}

func parent(tree map[string]string, last string) string {
	current := last
	for {
		parent, ok := tree[current]
		if !ok {
			break
		}
		current = parent
	}
	return current
}
