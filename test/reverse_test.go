package test

import (
	"testing"

	"echo-reverse-service/internal/app/reverse"
)

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := reverse.ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func ReverseRunes(s string) {
	panic("unimplemented")
}
