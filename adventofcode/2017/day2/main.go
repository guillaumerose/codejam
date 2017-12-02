package main

import (
	"strings"
	"strconv"
	"os"
	"bufio"
	"fmt"
	"math"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		min := math.MaxInt32
		max := math.MinInt32
		for _, i := range strings.Fields(scanner.Text()) {
			j, _ := strconv.Atoi(i)
			if min > j {
				min = j
			}
			if max < j {
				max = j
			}
		}
		sum += max - min
	}
	fmt.Println(sum)
}
