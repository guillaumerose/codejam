package main

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestFilter(t *testing.T) {
	assert.Equal(t, "{{}}", filter2("{{<!>},{<!>},{<!>},{<a>}}"))
	assert.Equal(t, "", filter2("<!!!>>"))
	assert.Equal(t, "{{},{},{},{}}", filter2("{{<!!>},{<!!>},{<!!>},{<!!>}}"))
}

func TestCount(t *testing.T) {
	assert.Equal(t, 1, count("{}"))
	assert.Equal(t, 3, count("{{}}"))
	assert.Equal(t, 5, count("{{},{}}"))
	assert.Equal(t, 6, count("{{{}}}"))
	assert.Equal(t, 9, count("{{},{},{},{}}"))
}

func TestSkipped(t *testing.T) {
	assert.Equal(t, 0, skipped("<>"))
	assert.Equal(t, 17, skipped("<random characters>"))
	assert.Equal(t, 3, skipped(`<<<<>`))
	assert.Equal(t, 2, skipped(`<{!>}>`))
	assert.Equal(t, 0, skipped(`<!!>`))
	assert.Equal(t, 0, skipped(`<!!!>>`))
	assert.Equal(t, 10, skipped(`<{o"i!a,<{i<a>`))
	assert.Equal(t, 0, skipped(`{{<!!>},{<!!>},{<!!>},{<!!>}}`))
}

func filter2(input string) string {
	out, _ := filter(input)
	return out
}

func skipped(input string) int {
	_, i := filter(input)
	return i
}
