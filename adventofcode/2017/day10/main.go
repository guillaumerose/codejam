package main

import (
	"encoding/hex"
)

func main() {
	input := make([]int, 256)
	for i := 0; i < 256; i++ {
		input[i] = i
	}
	lengths := []int{106, 118, 236, 1, 130, 0, 235, 254, 59, 205, 2, 87, 129, 25, 255, 118}
	reverse(lengths, input)
}

func hash(str string) string {
	input := make([]int, 256)
	for i := 0; i < 256; i++ {
		input[i] = i
	}
	lengths := []int{49, 44, 50, 44, 51, 17, 31, 73, 47, 23}
	if str == "" {
		lengths = []int{17, 31, 73, 47, 23}
	}
	len := len(input)
	skip := 0
	current := 0
	for r := 0; r < 64; r++ {
		for _, length := range lengths {
			for i := 0; i < length/2; i++ {
				x := (current + i) % len
				y := (current + length - 1 - i) % len
				a := input[x]
				b := input[y]
				input[x] = b
				input[y] = a
			}
			current = (current + length + skip) % len
			skip++
		}
	}
	acc := ""
	for z := 0; z < 16; z++ {
		xor := input[16*z]
		for k := 16*z + 1; k < 16*(z+1); k++ {
			xor ^= input[k]
		}
		acc += hex.EncodeToString([]byte{byte(xor)})
	}
	return acc
}

func reverse(lengths []int, input []int) {
	len := len(input)
	skip := 0
	current := 0
	for _, length := range lengths {
		for i := 0; i < length/2; i++ {
			x := (current + i) % len
			y := (current + length - 1 - i) % len
			a := input[x]
			b := input[y]
			input[x] = b
			input[y] = a
		}
		current = (current + length + skip) % len
		skip++
	}
}
