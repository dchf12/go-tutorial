package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		// TODO: Add test cases.
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tt := range tests {
		rev := Reverse(tt.in)
		if rev != tt.want {
			t.Errorf("Reverse: %q, want %q", rev, tt.want)
		}
	}
}
