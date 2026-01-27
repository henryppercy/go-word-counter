package e2e

import (
	"bytes"
	"testing"
)

func TestNoExist(t *testing.T) {
	cmd, err := getCommand("no_exist.txt")
	if err != nil {
		t.Fatal("failed to get pwd: ", err)
	}

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

	if stderr.String() != wantsStderr {
		t.Log("stderr is not correct: wanted: ", wantsStderr, "got: ", stderr.String())
		t.Fail()
	}

	if stdout.String() != wantsStdout {
		t.Log("stdout is not correct: wanted: ", wantsStdout, "got: ", stdout.String())
		t.Fail()
	}
}
