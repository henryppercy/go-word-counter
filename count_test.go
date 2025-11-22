package main_test

import (
	"strings"
	"testing"

	counter "github.com/henryppercy/counter"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "single space",
			input: " ",
			wants: 0,
		},
		{
			name:  "new lines",
			input: "one two three\nfour five",
			wants: 5,
		},
		{
			name:  "multiple spaces",
			input: "This is a sentence.  This is another.",
			wants: 7,
		},
		{
			name:  "prefixed multiple spaces",
			input: "   Hello",
			wants: 1,
		},
		{
			name:  "suffixed multiple spaces",
			input: "Hello   ",
			wants: 1,
		},
		{
			name:  "tabbed character",
			input: "Hello\tWorld\n",
			wants: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gives := counter.CountWords(strings.NewReader(tc.input))

			if gives != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, gives)
				t.Fail()
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	testCases := []struct{
		name  string
		input string
		wants int
	}{
		{
			name: "simple 5 words, 1 new line",
			input: "one two three four five\n",
			wants: 1,
		},
		{
			name: "empty file",
			input: "",
			wants: 0,
		},
		{
			name: "no new lines",
			input: "one two three four five",
			wants: 0,
		},
		{
			name: "no new lines at end",
			input: "one two three four five\nsix",
			wants: 1,
		},
		{
			name: "multi newline input",
			input: "\n\n\n\n",
			wants: 4,
		},
		{
			name: "multi word line string",
			input: "one\ntwo\nthree\nfour\nfive\n",
			wants: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gives := counter.CountLines(strings.NewReader(tc.input))

			if gives != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, gives)
				t.Fail()
			}
		})
	}
}
