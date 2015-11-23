package spec

import (
	"testing"
)

func TestIsValidName(t *testing.T) {
	cases := []struct {
		testCase string
		expect   bool
	}{
		{"test", true},
		{"test1", true},
		{"test-1", true},
		{"test_1", true},
		{"Test", true},
		{"test.1", false},
		{"test%", false},
		{"test/", false},
	}

	for _, c := range cases {
		got := isValidName(c.testCase)
		if c.expect != got {
			t.Errorf("For filename %q, Expect %t, got %t", c.testCase, c.expect, got)
		}
	}

}

//func TestValidateMetadata(t *testing.T) {
//	cases := []struct {
//		testCase string
//		expect   bool
//	}{
//		{"test", true},
//		{"test1", true},
//		{"test-1", true},
//		{"test_1", true},
//		{"Test", true},
//		{"test.1", false},
//		{"test%", false},
//		{"test/", false},
//	}
//
//	for _, c := range cases {
//		got := isValidName(c.testCase)
//		if c.expect != got {
//			t.Errorf("For filename %q, Expect %t, got %t", c.testCase, c.expect, got)
//		}
//	}
//
//}
