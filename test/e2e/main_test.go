package e2e

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var binName = "counter_test"

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", binName, "../../cmd")

	var buf bytes.Buffer
	cmd.Stderr = &buf

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to build binary:", err)
		fmt.Fprintln(os.Stderr, buf.String())
		os.Exit(1)
	}

	code := m.Run()

	_ = os.Remove(binName)

	os.Exit(code)
}

func getCommand(args ...string) (*exec.Cmd, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, binName)

	cmd := exec.Command(path, args...)

	return cmd, nil
}

func createFile(content string) (*os.File, error) {
	file, err := os.CreateTemp("", "counter-test-*")
	if err != nil {
		return nil, fmt.Errorf("create a file: %w", err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		return nil, fmt.Errorf("failed to write content: %w", err)
	}

	err = file.Close()
	if err != nil {
		return nil, fmt.Errorf("close file: %w", err)
	}

	return file, nil
}
