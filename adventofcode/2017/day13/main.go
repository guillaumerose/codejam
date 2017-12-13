package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var fw map[int]int
var max int

func main() {
	fw = make(map[int]int)
	sign := map[int]int{}
	state := map[int]int{}
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

	fmt.Println(find(state, sign))
}

func find(state map[int]int, sign map[int]int) int {
	i := 0
	packets := map[int]int{}
	for {
		packets[i] = 0
		for start, pos := range packets {
			if _, ok := fw[pos]; state[pos] == 0 && ok {
				delete(packets, start)
			} else {
				packets[start]++
				if packets[start] > max {
					return start
				}
			}
		}
		move(state, sign)
		i++
	}
}

func move(state map[int]int, sign map[int]int) {
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
