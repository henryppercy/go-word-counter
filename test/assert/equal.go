package assert

import (
	"reflect"
	"strings"
	"testing"
)

func Equal(t *testing.T, wants any, got any, msgs ...string) {
	t.Helper()

	if reflect.DeepEqual(wants, got) {
		return
	}

	msgStr := strings.Join(msgs, ": ")

	t.Logf("%s\nexpected: %v\n     got: %v\n", msgStr, wants, got)
	t.Fail()
}

func Error(t *testing.T, err error, msgs ...string) {
	t.Helper()

	if err == nil {
		return
	}

	msgStr := strings.Join(msgs, ": ")

	t.Fatal(msgStr, err)
	// t.Fail()
}
