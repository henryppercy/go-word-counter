package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	wants := 5
	gives := countWords([]byte(input))

	if gives != wants {
		t.Fail()
	}
}
