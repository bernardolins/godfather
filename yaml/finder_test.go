package yaml

import (
	"testing"
)

func TestIsYaml(t *testing.T) {
	cases := []struct {
		filename string
		expect   bool
	}{
		{".godfather.yaml", true},
		{".godfather.yml", true},
		{".godfather", false},
		{"godfather.yaml", true},
		{"fathergod.yml", true},
	}

	for _, c := range cases {
		got := isYaml(c.filename)
		if c.expect != got {
			t.Errorf("For filename %q, Expect %t, got %t", c.filename, c.expect, got)
		}
	}

}

func TestIsGodfatherFile(t *testing.T) {
	cases := []struct {
		filename string
		expect   bool
	}{
		{".godfather.yaml", true},
		{".godfather.yml", true},
		{".godfather", false},
		{"godfather.yaml", false},
		{"fathergod.yml", false},
	}

	for _, c := range cases {
		got := isGodfatherFile(c.filename)
		if c.expect != got {
			t.Errorf("For filename %q, Expect %t, got %t", c.filename, c.expect, got)
		}
	}

}
