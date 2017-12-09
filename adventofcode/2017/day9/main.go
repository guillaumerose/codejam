package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bin, _ := ioutil.ReadAll(os.Stdin)
	input := string(bin)
	_, skipped := filter(input)
	fmt.Println(skipped)
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

func filter(input string) (string, int) {
	output := ""
	i := 0
	skip := 0
	skipped := 0
	for i < len(input) {
		current := string(input[i])
		// fmt.Println(i, current)
		if current == "!" {
			i += 2
			continue
		}
		if current == "<" {
			skip += 1
		}
		if skip == 0 {
			output += current
		}
		if skip > 0 && (skip != 1 || current != "<") && current != ">" {
			skipped++
		}
		if current == ">" {
			skip = 0
		}
		i++
	}
	return output, skipped
}
