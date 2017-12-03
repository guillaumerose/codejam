package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := pos(368078)
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func pos(i int) (int, int) {
	size := int(math.Ceil(math.Sqrt(float64(i))))
	if size%2 == 0 {
		size++
	}
	x := size / 2
	y := -size / 2
	max := size * size
	int1 := max - (size - 1)
	int2 := int1 - (size - 1)
	int3 := int2 - (size - 1)
	int4 := int3 - (size - 1)
	if max >= i && i > int1 {
		return x - (max - i), y
	}
	if int1 >= i && i > int2 {
		return -x, y + (int1 - i)
	}
	if int2 >= i && i > int3 {
		return x + (int3 - i), -y
	}
	return x, y - (int4 - i)
}
