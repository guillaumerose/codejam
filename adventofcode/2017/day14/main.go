package main

import (
	"fmt"
	"gopkg.in/fatih/set.v0"
)

type node struct {
	x, y int
}

func main() {
	input := "stpzcrnm"
	matrix := make([][]int, 128)
	for i := 0; i < 128; i++ {
		lengths := getInputLengths(fmt.Sprintf("%s-%d", input, i))
		hashList := generateList(256)
		doTheHash(hashList, lengths, 64)
		denseHash := convertToDenseHash(hashList)
		hex := convertToHex(denseHash)

		matrix[i] = make([]int, 128)
		for j, char := range hex {
			if char == 49 {
				matrix[i][j] = -1
			} else {
				matrix[i][j] = 0
			}
		}
		fmt.Println(matrix[i])
	}

	fmt.Println("---")

	color := 1
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if matrix[i][j] == -1 {
				set := set.New(node{i, j})
				for !set.IsEmpty() {
					val := set.Pop().(node)
					fmt.Println(val)
					x, y := val.x, val.y
					if x < 0 || x > 127 {
						continue
					}
					if y < 0 || y > 127 {
						continue
					}
					if matrix[x][y] != -1 {
						continue
					}
					matrix[x][y] = color
					set.Add(node{x - 1, y})
					set.Add(node{x + 1, y})
					set.Add(node{x, y - 1})
					set.Add(node{x, y + 1})
				}
				color++
			}
		}
	}

	for i := 0; i < 128; i++ {
		fmt.Println(matrix[i])
	}

	fmt.Println(color - 1)
}

func mod(number int, length int) int {
	if number < 0 {
		return mod(length+number, length)
	} else {
		return number % length
	}
}

func revertSubList(list []int, start int, stop int, length int) {
	copyList := make([]int, len(list))
	copy(copyList, list)
	for i := 0; i < length; i++ {
		index1 := mod(start+i, len(list))
		index2 := mod(stop-i, len(list))
		list[index1] = copyList[index2]
	}
}

func doTheHash(list []int, inputLengths []int, numberRounds int) {
	start := 0
	skipSize := 0
	for i := 0; i < numberRounds; i++ {
		for _, length := range inputLengths {
			startInvert := (start) % len(list)
			endInvert := (start + length - 1) % len(list)
			revertSubList(list, startInvert, endInvert, length)
			start += length + skipSize
			skipSize++
		}
	}

}

func generateList(length int) []int {
	intList := make([]int, length)
	for i := 0; i < length; i++ {
		intList[i] = i
	}
	return intList
}
func getInputLengths(input string) []int {
	lengths := append([]byte(input), 17, 31, 73, 47, 23)
	intLengths := make([]int, len(lengths))
	for i, element := range lengths {
		intLengths[i] = int(element)
	}
	return intLengths
}

func convertToDenseHash(list []int) []int {
	denseHash := make([]int, len(list)/16)
	for i := 0; i < len(list); i += 16 {
		for j := 0; j < 16; j++ {
			denseHash[i/16] ^= list[i+j]
		}
	}
	return denseHash
}

func convertToHex(list []int) string {
	hexString := ""
	for _, element := range list {
		val := fmt.Sprintf("%08b", element)
		hexString += val
	}
	return hexString
}
