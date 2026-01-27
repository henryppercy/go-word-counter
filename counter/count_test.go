package counter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/henryppercy/counter/display"
	"github.com/henryppercy/counter/test/assert"
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
			r := strings.NewReader(tc.input)
			gives := GetCount(r).words

			assert.Equal(t, tc.wants, gives)
		})
	}
}

func TestCountLines(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "simple 5 words, 1 new line",
			input: "one two three four five\n",
			wants: 1,
		},
		{
			name:  "empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "no new lines",
			input: "one two three four five",
			wants: 0,
		},
		{
			name:  "no new lines at end",
			input: "one two three four five\nsix",
			wants: 1,
		},
		{
			name:  "multi newline input",
			input: "\n\n\n\n",
			wants: 4,
		},
		{
			name:  "multi word line string",
			input: "one\ntwo\nthree\nfour\nfive\n",
			wants: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			gives := GetCount(r).lines

			assert.Equal(t, tc.wants, gives)
		})
	}
}

func TestCountBytes(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "five words",
			input: "one two three four five",
			wants: 23,
		},
		{
			name:  "empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "all spaces",
			input: "       ",
			wants: 7,
		},
		{
			name:  "newlines, tabs, and words",
			input: "one\ntwo\nthree\nfour\t\n",
			wants: 20,
		},
		{
			name:  "Unicode characters",
			input: "Mañana",
			wants: 7,
		},
		{
			name:  "",
			input: "",
			wants: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			gives := GetCount(r).bytes

			assert.Equal(t, tc.wants, gives)
		})
	}
}

func TestGetCounts(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants Counts
	}{
		{
			name:  "simple five words",
			input: "one two three four five\n",
			wants: Counts{
				lines: 1,
				words: 5,
				bytes: 24,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			gives := GetCount(r)

			assert.Equal(t, tc.wants, gives)
		})
	}
}

func TestPrintCount(t *testing.T) {
	type inputs struct {
		counts   Counts
		opts     display.NewOptionsArgs
		filename []string
	}
	testCases := []struct {
		name  string
		input inputs
		wants string
	}{
		{
			name: "simple five words",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				opts: display.NewOptionsArgs{
					ShowBytes: true,
					ShowLines: true,
					ShowWords: true,
				},
				filename: []string{"words.txt"},
			},
			wants: "1\t5\t24\t words.txt\n",
		},
		{
			name: "no file name",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 4,
					bytes: 20,
				},
				opts: display.NewOptionsArgs{
					ShowBytes: true,
					ShowWords: true,
					ShowLines: true,
				},
			},
			wants: "1\t4\t20\t\n",
		},
		{
			name: "simple five words no options",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				filename: []string{"words.txt"},
			},
			wants: "1\t5\t24\t words.txt\n",
		},
		{
			name: "simple five words show lines",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				opts: display.NewOptionsArgs{
					ShowBytes: false,
					ShowWords: false,
					ShowLines: true,
				},
				filename: []string{"words.txt"},
			},
			wants: "1\t words.txt\n",
		},
		{
			name: "simple five words show words",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				opts: display.NewOptionsArgs{
					ShowBytes: false,
					ShowWords: true,
					ShowLines: false,
				},
				filename: []string{"words.txt"},
			},
			wants: "5\t words.txt\n",
		},
		{
			name: "simple five words show bytes and lines",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				opts: display.NewOptionsArgs{
					ShowBytes: true,
					ShowWords: false,
					ShowLines: true,
				},
				filename: []string{"words.txt"},
			},
			wants: "1\t24\t words.txt\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			tc.input.counts.Print(buffer, display.NewOptions(tc.input.opts), tc.input.filename...)

			assert.Equal(t, tc.wants, buffer.String())
		})
	}
}

func TestAddCount(t *testing.T) {
	type inputs struct {
		counts Counts
		other  Counts
	}
	testCases := []struct {
		name  string
		input inputs
		wants Counts
	}{
		{
			name: "simple add by one",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				other: Counts{
					lines: 1,
					words: 1,
					bytes: 1,
				},
			},
			wants: Counts{
				lines: 2,
				words: 6,
				bytes: 25,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			totals := tc.input.counts
			gives := totals.Add(tc.input.other)

			assert.Equal(t, tc.wants, gives)
		})
	}
}

var benchData = []string{
	"This is a test data string\nthat spans across\nmultiple lines\n",
	"one two three\nfour five\nsix\nseven\neight\n",
	"this is a weird\n\n\n\n\n\n\n        string\n",
}

func BenchmarkGetCount(b *testing.B) {
	for i := range b.N {
		data := benchData[i%len(benchData)]

		r := strings.NewReader(data)

		GetCount(r)
	}
}

func BenchmarkGetCountSinglePass(b *testing.B) {
	for i := range b.N {
		data := benchData[i%len(benchData)]

		r := strings.NewReader(data)

		GetCountSinglePass(r)
	}
}
