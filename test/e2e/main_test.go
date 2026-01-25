package e2e

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
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
