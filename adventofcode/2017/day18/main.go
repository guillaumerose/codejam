package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var prog []string
	for scanner.Scan() {
		prog = append(prog, scanner.Text())
	}
	a := make(chan int, 1000000000)
	b := make(chan int, 1000000000)
	stop1 := make(chan int)
	stop2 := make(chan int)
	go run(prog, 0, a, b, stop1)
	run(prog, 1, b, a, stop2)
	close(a)
	close(b)
	<-stop1
	<-stop2
}

func run(prog []string, p int, in chan int, out chan int, stop chan int) {
	reg := make(map[string]int)
	reg["p"] = p
	ins := 0
	for ins < len(prog) {
		text := prog[ins]
		split := strings.Fields(text)
		op := split[0]
		//fmt.Printf("%d: %s\n", p, text)
		switch op {
		case "set":
			y := split[2]
			i, err := strconv.Atoi(y)
			x := split[1]
			if err != nil {
				reg[x] = reg[y]
			} else {
				reg[x] = i
			}
		case "add":
			y := split[2]
			i, err := strconv.Atoi(y)
			x := split[1]
			if err != nil {
				reg[x] += reg[y]
			} else {
				reg[x] += i
			}
		case "mul":
			y := split[2]
			i, err := strconv.Atoi(y)
			x := split[1]
			if err != nil {
				reg[x] *= reg[y]
			} else {
				reg[x] *= i
			}
		case "mod":
			y := split[2]
			i, err := strconv.Atoi(y)
			x := split[1]
			if err != nil {
				reg[x] %= reg[y]
			} else {
				reg[x] %= i
			}
		case "rcv":
			//	fmt.Printf("%d: in size %d\n", p, len(in))
			reg[split[1]] = <-in
			fmt.Printf("%d: rcv %d\n", p, reg[split[1]])
		case "snd":
			fmt.Printf("%d: snd %d\n", p, reg[split[1]])
			out <- reg[split[1]]
			//fmt.Printf("%d: out size %d\n", p, len(out))
		case "jgz":
			i, err := strconv.Atoi(split[1])
			if err != nil {
				i = reg[split[1]]
			}
			if i > 0 {
				jmp, err := strconv.Atoi(split[2])
				if err != nil {
					jmp = reg[split[2]]
				}
				ins += jmp - 1
			}
		}
		//fmt.Printf("%d: %v\n", p, reg)
		ins++
	}
	stop <- 1
}
