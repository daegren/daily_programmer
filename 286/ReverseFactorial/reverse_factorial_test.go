package ReverseFactorial

import "testing"

func TestReverseFactiorial(t *testing.T) {
	var cases = []struct {
		in   int
		want string
	}{
		{120, "5!"},
		{150, "NONE"},
		{24, "4!"},
		{123, "NONE"},
	}
	for _, c := range cases {
		got := IsFactorial(c.in)
		if got != c.want {
			t.Errorf("IsFactorial(%d) == (%s), want %s", c.in, got, c.want)
		}
	}
}
