package main

import "fmt"

const pairs = 40000000
const factorA = 16807
const factorB = 48271

func main() {
	seedA := 873
	seedB := 583
	count := 0
	for i := 0; i < pairs; i++ {
		seedA = (seedA * factorA) % 2147483647
		seedB = (seedB * factorB) % 2147483647

		if seedA&0xFFFF == seedB&0xFFFF {
			count++
		}
	}
	fmt.Println(count)
}
