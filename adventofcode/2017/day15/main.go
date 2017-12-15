package main

import "fmt"

const pairs = 5000000
const factorA = 16807
const factorB = 48271
const mod = 2147483647

func main() {
	seedA := 873
	seedB := 583
	count := 0
	for i := 0; i < pairs; i++ {
		seedA = (seedA * factorA) % mod
		for seedA%4 != 0 {
			seedA = (seedA * factorA) % mod
		}
		seedB = (seedB * factorB) % mod
		for seedB%8 != 0 {
			seedB = (seedB * factorB) % mod
		}

		if seedA&0xFFFF == seedB&0xFFFF {
			count++
		}
	}
	fmt.Println(count)
}
