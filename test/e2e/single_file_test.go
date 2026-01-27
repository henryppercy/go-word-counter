package e2e

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestSingleFile(t *testing.T) {
	file, err := os.CreateTemp("", "counter-test-*")
	if err != nil {
		t.Fatal("failed to create temp file:", err)
	}

	defer os.Remove(file.Name())

	_, err = file.WriteString("foo bar baz\nbaz bar foo\none two three\n")
	if err != nil {
		t.Fatal("failed to write to temp file", err)
	}

	err = file.Close()
	if err != nil {
		t.Fatal("failed to close file:", err)
	}

	cmd, err := getCommand(file.Name())
	if err != nil {
		t.Fatal("failed to get pwd: ", err)
	}

	output := &bytes.Buffer{}
	cmd.Stdout = output

	if err = cmd.Run(); err != nil {
		t.Fatal("failed to run command", err)
	}

	wants := fmt.Sprintf(" 3 9 38 %s\n", file.Name())
	if output.String() != wants {
		t.Log("stdout is not correct: wanted: ", wants, "got: ", output.String())
		t.Fail()
	}
}
