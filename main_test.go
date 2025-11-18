package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	wants := 5
	gives := CountWords([]byte(input))

	if gives != wants {
		t.Logf("expected: %d got: %d", wants, gives)
		t.Fail()
	}

	input = ""
	wants = 0
	gives = CountWords([]byte(input))

	if gives != wants {
		t.Logf("expected: %d got: %d", wants, gives)
		t.Fail()
	}

	input = " "
	wants = 0
	gives = CountWords([]byte(input))

	if gives != wants {
		t.Logf("expected: %d got: %d", wants, gives)
		t.Fail()
	}
}
