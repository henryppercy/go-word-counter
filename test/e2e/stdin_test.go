package e2e

import (
	"bytes"
	"strings"
	"testing"

	"github.com/henryppercy/counter/test/assert"
)

func TestStdin(t *testing.T) {
	cmd, err := getCommand()
	assert.Error(t, err, "failed to get command")

	output := &bytes.Buffer{}

	cmd.Stdin = strings.NewReader("one two three\n")
	cmd.Stdout = output

	err = cmd.Run()
	assert.Error(t, err, "failed to run command")

	wants := " 1 3 14\n"
	assert.Equal(t, wants, output.String(), "stdout is not valid")
}
