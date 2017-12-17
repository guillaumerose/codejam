package main

import "fmt"

func main() {
	step := 304
	fmt.Println(step)
	var t []int
	t = append(t, 0)
	current := 0
	for i := 1; i <= 50000000; i++ {
		current = (current + step) % i
		if current+1 == 1 {
			fmt.Println(i)
		}
		current++
	}
}
