package main

func main() {
	input := make([]int, 256)
	for i := 0; i < 256; i++ {
		input[i] = i
	}
	lengths := []int{106, 118, 236, 1, 130, 0, 235, 254, 59, 205, 2, 87, 129, 25, 255, 118}
	reverse(lengths, input)
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
