package e2e

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/henryppercy/counter/test/assert"
)

func TestSingleFile(t *testing.T) {
	file, err := os.CreateTemp("", "counter-test-*")
	assert.Error(t, err, "failed to create temp file")

	defer os.Remove(file.Name())

	_, err = file.WriteString("foo bar baz\nbaz bar foo\none two three\n")
	assert.Error(t, err, "failed to write temp file")

	err = file.Close()
	assert.Error(t, err, "failed to close file")

	cmd, err := getCommand(file.Name())
	assert.Error(t, err, "failed to get command")

	output := &bytes.Buffer{}
	cmd.Stdout = output

	err = cmd.Run()
	assert.Error(t, err, "failed to run command")

	wants := fmt.Sprintf(" 3 9 38 %s\n", file.Name())

	assert.Equal(t, wants, output.String())
}
