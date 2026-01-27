package e2e

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestFlags(t *testing.T) {
	file, err := createFile("one two three four five\none two thee\n")
	if err != nil {
		t.Fatal("failed to create file:", err)
	}
	defer os.Remove(file.Name())

	testCases := []struct {
		name  string
		flags []string
		wants string
	}{
		{
			name:  "line flag",
			flags: []string{"-l"},
			wants: fmt.Sprintf(" 2 %s\n", file.Name()),
		},
		{
			name:  "bytes flag",
			flags: []string{"-c"},
			wants: fmt.Sprintf(" 37 %s\n", file.Name()),
		},
		{
			name:  "words flag",
			flags: []string{"-w"},
			wants: fmt.Sprintf(" 8 %s\n", file.Name()),
		},
	}

	for _, tc := range testCases {
		t.Run("lines flag", func(t *testing.T) {
			inputs := append(tc.flags, file.Name())

			cmd, err := getCommand(inputs...)
			if err != nil {
				t.Log("failed to get command:", err)
				t.Fail()
			}

			stdout := &bytes.Buffer{}
			cmd.Stdout = stdout

			if err = cmd.Run(); err != nil {
				t.Log("failed to run command:", err)
				t.Fail()
			}

			if tc.wants != stdout.String() {
				t.Logf("expected: %s got: %s", tc.wants, stdout.String())
				t.Fail()
			}
		})
	}
}
