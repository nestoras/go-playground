package git

import (
	"os/exec"
	"testing"
)

func TestVersion(t *testing.T) {
	execCommand = func(name string, arg ...string) *exec.Cmd {
		return exec.Command("echo", "git version 3.0.0")
	}
	defer func() {
		execCommand = exec.Command
	}()

	got := Version()
	want := "3.0.0"
	if got != want {
		t.Errorf("Version() = %q; want %q", got, want)
	}
}

func TestChecker_Version(t *testing.T) {
	checker := Checker{
		execCommand: func(name string, arg ...string) *exec.Cmd {
			return exec.Command("echo", "git version 3.0.0")
		},
	}
	got := checker.Version()
	want := "3.0.0"
	if got != want {
		t.Errorf("checker.Version() = %q; want %q", got, want)
	}
}