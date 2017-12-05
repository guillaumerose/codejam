package main

import "testing"

func TestProcess(t *testing.T) {
	assert(t, []int{0, 3, 0, 1, -3}, 10)
}

func assert(t *testing.T, input []int, expected int) {
	actual := process(input)
	if actual != expected {
		t.Fatalf("process(%v) = %d, got %d", input, expected, actual)
	}
}
