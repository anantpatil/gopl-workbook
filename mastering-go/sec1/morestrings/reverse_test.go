package morestrings

import "testing"

func TestReverseString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, World!", "!dlroW ,olleH"},
		{"123", "321"},
		{"", ""},
	}

	for _, c := range cases {
		got := ReverseString(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
