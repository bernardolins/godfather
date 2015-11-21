package config

import (
	"os"
	"testing"
)

func TestGetGodfatherPath(t *testing.T) {
	os.Setenv("GODFATHER_PATH", "~")
	expect := "~"
	got := GetGodfatherPath()

	if expect != got {
		t.Errorf("Expect %q, but got %q", expect, got)
	}
}

func TestGetGodfatherPathNotSet(t *testing.T) {
	os.Unsetenv("GODFATHER_PATH")
	expect := ""
	got := GetGodfatherPath()

	if expect != got {
		t.Errorf("Expect %q, but got %q", expect, got)
	}
}
