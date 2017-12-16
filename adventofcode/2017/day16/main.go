package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := "abcdefghijklmnop"
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")
		fmt.Println(split)
		for _, current := range split {
			switch string(current[0]) {
			case "s":
				i, _ := strconv.Atoi(current[1:])
				input = spin(input, i)
			case "x":
				t := strings.Split(current[1:], "/")
				x, _ := strconv.Atoi(t[0])
				y, _ := strconv.Atoi(t[1])
				input = exch(input, x, y)
			case "p":
				t := strings.Split(current[1:], "/")
				input = partner(input, t[0], t[1])
			}
		}
	}
	fmt.Println(input)
}

func partner(input string, a string, b string) string {
	input = strings.Replace(input, a, "z", 1)
	input = strings.Replace(input, b, a, 1)
	input = strings.Replace(input, "z", b, 1)
	return input
}

func exch(input string, y int, x int) string {
	tmp := input[y]
	input = replaceAtIndex(input, input[x], y)
	input = replaceAtIndex(input, tmp, x)
	return input
}

func spin(input string, spin int) string {
	input = input[len(input)-spin:len(input)] + input[0:len(input)-spin]
	return input
}

func replaceAtIndex(in string, r uint8, i int) string {
	out := []rune(in)
	out[i] = rune(r)
	return string(out)
}
