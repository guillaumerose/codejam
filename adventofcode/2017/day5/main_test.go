package main

import "testing"

func TestProcess(t *testing.T) {
	assert(t, "1212", 6)
	assert(t, "1221", 0)
	assert(t, "123425", 4)
	assert(t, "123123", 12)
	assert(t, "12131415", 4)
}

func assert(t *testing.T, input string, expected int) {
	actual := process(input)
	if actual != expected {
		t.Fatalf("process(%s) = %d, got %d", input, expected, actual)
	}
}
