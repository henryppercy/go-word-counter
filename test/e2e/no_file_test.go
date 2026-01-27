package e2e

import (
	"bytes"
	"testing"

	"github.com/henryppercy/counter/test/assert"
)

func TestNoExist(t *testing.T) {
	cmd, err := getCommand("no_exist.txt")
	assert.Error(t, err, "failed to get command")

	stderr := &bytes.Buffer{}
	stdout := &bytes.Buffer{}

	cmd.Stderr = stderr
	cmd.Stdout = stdout

	wantsStderr := "counter: open no_exist.txt: no such file or directory\n"
	wantsStdout := ""

	if err = cmd.Run(); err == nil {
		t.Log("command succeeded when should have failed")
		t.Fail()
	}

	if err.Error() != "exit status 1" {
		t.Log("expected error of exit status 1. got:", err.Error())
		t.Fail()
	}

	assert.Equal(t, wantsStderr, stderr.String(), "stderr doesn't match")
	assert.Equal(t, wantsStdout, stdout.String(), "stdout doesn't match")
}
