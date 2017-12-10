package main

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestFilter(t *testing.T) {
	input := []int{0, 1, 2, 3, 4}
	reverse([]int{3, 4, 1, 5}, input)
	assert.Equal(t, []int{3, 4, 2, 1, 0}, input)
}
