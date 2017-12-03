package main

import "testing"

func TestProcess(t *testing.T) {
	assert(t, 1, 0, 0)
	assert(t, 2, 1, 0)
	assert(t, 3, 1, 1)
	assert(t, 4, 0, 1)
	assert(t, 5, -1, 1)
	assert(t, 6, -1, 0)
	assert(t, 7, -1, -1)
	assert(t, 8, 0, -1)
	assert(t, 9, 1, -1)
	assert(t, 10, 2, -1)
}

func assert(t *testing.T, input, x, y int) {
	a, b := pos(input)
	if a != x || b != y {
		t.Fatalf("pos(%d) = (%d, %d), got (%d, %d)", input, x, y, a, b)
	}
}
