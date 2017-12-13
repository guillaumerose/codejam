package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	fw := map[int]int{}
	sign := map[int]int{}
	state := map[int]int{}
	max := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ": ")
		k, _ := strconv.Atoi(split[0])
		v, _ := strconv.Atoi(split[1])
		fw[k] = v
		sign[k] = 1
		state[k] = 0
		if k > max {
			max = k
		}
	}

	caught := 0
	for current := 0; current <= max; current++ {
		if _, ok := fw[current]; state[current] == 0 && ok {
			fmt.Printf("Caught: %d\n", current)
			caught += fw[current] * current
		}
		move(state, sign, fw)
	}
	fmt.Println(caught)
}

func move(state map[int]int, sign map[int]int, fw map[int]int) {
	for i, _ := range state {
		if state[i] == 0 {
			sign[i] = 1
		}
		if state[i] == fw[i]-1 {
			sign[i] = -1
		}
		state[i] = (state[i] + sign[i]*1)
	}
}
