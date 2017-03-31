package stringutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"世界 你好", "好你 界世"},
		{"Hello World", "dlrow olleH"},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q , want %q", c.in, got, c.want)
		}
	}
}
