package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bin, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(count(filter(string(bin))))
}

func count(input string) int {
	score := 0
	level := 1
	for i := 0; i < len(input); i++ {
		current := string(input[i])
		if current == "{" {
			score += level
			level++
		}
		if current == "}" {
			level--
		}
	}
	return score
}

func filter(input string) string {
	output := ""
	i := 0
	skip := false
	for i < len(input) {
		current := string(input[i])
		// fmt.Println(i, current)
		if current == "!" {
			i += 2
			continue
		}
		if current == "<" {
			skip = true
		}
		if !skip {
			output += current
		}
		if current == ">" {
			skip = false
		}
		i++
	}
	return output
}
