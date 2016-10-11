package Kaprekar

import "testing"

func TestLargestDigit(t *testing.T) {
	var cases = []struct {
		in   int
		want int
	}{
		{1234, 4},
		{3253, 5},
		{9800, 9},
		{3333, 3},
		{120, 2},
	}
	for _, c := range cases {
		got := LargestDigit(c.in)
		if got != c.want {
			t.Errorf("LargestDigit(%d) == (%d), want %d", c.in, got, c.want)
		}
	}
}
