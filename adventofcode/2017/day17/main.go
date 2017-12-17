package main

import "fmt"

func main() {
	step := 304
	fmt.Println(step)
	var t []int
	t = append(t, 0)
	current := 0
	for i := 1; i <= 2017; i++ {
		current = (current + step) % len(t)
		t = append(t[:current+1], append([]int{i}, t[current+1:]...)...)
		current++
		fmt.Println(current, t)
	}
	fmt.Println(t[current+1])
}
