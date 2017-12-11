package main

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestProcess(t *testing.T) {
	assert.Equal(t, map[string]int{}, reduce([]string{"n", "s"}))
	assert.Equal(t, map[string]int{"n": 2}, reduce([]string{"n", "n", "n", "s"}))
	assert.Equal(t, map[string]int{"s": 2}, reduce([]string{"n", "s", "s", "s"}))
	assert.Equal(t, map[string]int{}, reduce([]string{"ne", "sw"}))
	assert.Equal(t, map[string]int{"s": 2, "sw": 1}, reduce([]string{"se", "sw", "se", "sw", "sw"}))
	assert.Equal(t, map[string]int{"se": 2}, reduce([]string{"ne", "ne", "s", "s"}))
}
