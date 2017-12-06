package main

import "fmt"

func main() {
	seen := make(map[string]int)
	input := []int{4, 10, 4, 1, 8, 4, 9, 14, 5, 1, 14, 15, 0, 15, 3, 5}

	i := 0
	for {
		cycle(input)
		i++
		if first, ok := seen[hash(input)]; ok {
			fmt.Println(i - first)
			break
		}
		seen[hash(input)] = i
	}
}

func hash(input []int) string {
	return fmt.Sprintf("%v", input)
}

func cycle(input []int) {
	val, index := max(input)
	input[index] = 0
	for j := 1; j <= val; j++ {
		input[(index+j)%len(input)] += 1
	}
}

func max(input []int) (int, int) {
	index := -1
	max := -1
	for i := 0; i < len(input); i++ {
		if max < input[i] {
			max = input[i]
			index = i
		}
	}
	return max, index
}
