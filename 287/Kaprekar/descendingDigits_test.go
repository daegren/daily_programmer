package Kaprekar

import "testing"

func TestDescendingDigits(t *testing.T) {
	var cases = []struct {
		in   int
		want int
	}{
		{1234, 4321},
		{3253, 5332},
		// {9800, 9800},
		// {3333, 3333},
		// {120, 2100},
	}
	for _, c := range cases {
		got := DescendingDigits(c.in)
		if got != c.want {
			t.Errorf("DescendingDigits(%d) == (%d), want %d", c.in, got, c.want)
		}
	}
}
