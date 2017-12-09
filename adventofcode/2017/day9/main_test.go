package main

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestFilter(t *testing.T) {
	assert.Equal(t, "{{}}", filter("{{<!>},{<!>},{<!>},{<a>}}"))
	assert.Equal(t, "", filter("<!!!>>"))
	assert.Equal(t, "{{},{},{},{}}", filter("{{<!!>},{<!!>},{<!!>},{<!!>}}"))
}

func TestCount(t *testing.T) {
	assert.Equal(t, 1, count("{}"))
	assert.Equal(t, 3, count("{{}}"))
	assert.Equal(t, 5, count("{{},{}}"))
	assert.Equal(t, 6, count("{{{}}}"))
	assert.Equal(t, 9, count("{{},{},{},{}}"))
}
