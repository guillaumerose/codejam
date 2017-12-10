package main

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestReverse(t *testing.T) {
	input := []int{0, 1, 2, 3, 4}
	reverse([]int{3, 4, 1, 5}, input)
	assert.Equal(t, []int{3, 4, 2, 1, 0}, input)
}

func TestHash(t *testing.T) {
	assert.Equal(t, "3efbe78a8d82f29979031a4aa0b16a9d", hash("1,2,3"))
	assert.Equal(t, "a2582a3a0e66e6e86e3812dcb672a272", hash(""))
}
