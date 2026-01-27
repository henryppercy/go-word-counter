package e2e

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/henryppercy/counter/test/assert"
)

func TestMultipleFilesDeterministic(t *testing.T) {
	fileA, err := createFile("one two three four five\n")
	assert.Error(t, err, "failed to create fileA")

	defer os.Remove(fileA.Name())

	fileB, err := createFile("foo bar baz\n\n")
	assert.Error(t, err, "failed to create fileB")

	defer os.Remove(fileB.Name())

	fileC, err := createFile("")
	assert.Error(t, err, "failed to create fileC")

	defer os.Remove(fileC.Name())

	cmd, err := getCommand(fileA.Name(), fileB.Name(), fileC.Name())
	assert.Error(t, err, "failed to create command")

	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	err = cmd.Run()
	assert.Error(t, err, "failed to run command")

	wants := fmt.Sprintf(` 1 5 24 %s
 2 3 13 %s
 0 0  0 %s
 3 8 37 total
`, fileA.Name(), fileB.Name(), fileC.Name())

	assert.Equal(t, wants, stdout.String())
}

func TestMultipleFilesNonDeterministic(t *testing.T) {
	fileA, err := createFile("one two three four five\n")
	assert.Error(t, err, "failed to create fileA")

	defer os.Remove(fileA.Name())

	fileB, err := createFile("foo bar baz\n\n")
	assert.Error(t, err, "failed to create fileB")

	defer os.Remove(fileB.Name())

	fileC, err := createFile("")
	assert.Error(t, err, "failed to create fileC")

	defer os.Remove(fileC.Name())

	cmd, err := getCommand(fileA.Name(), fileB.Name(), fileC.Name())
	assert.Error(t, err, "failed to get command")

	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	err = cmd.Run()
	assert.Error(t, err, "failed to run command")

	wants := map[string]string{
		fileA.Name(): fmt.Sprintf(" 1 5 24 %s", fileA.Name()),
		fileB.Name(): fmt.Sprintf(" 2 3 13 %s", fileB.Name()),
		fileC.Name(): fmt.Sprintf(" 0 0  0 %s", fileC.Name()),
		"total":      " 3 8 37 total",
	}

	scanner := bufio.NewScanner(stdout)
	checkedWants := 0

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) == 0 {
			t.Log("line was empty")
			t.Fail()
		}

		filename := fields[len(fields)-1]

		lineWants, ok := wants[filename]
		if !ok {
			t.Logf("no wants for filename: %s", filename)
			t.Fail()
			continue
		}

		checkedWants++

		assert.Equal(t, lineWants, line)
	}

	if checkedWants != len(wants) {
		t.Logf("only checked: %d expected to check: %d", checkedWants, len(wants))
		t.Fail()
	}
}
