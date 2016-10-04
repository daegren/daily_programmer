package ReverseFactorial

import "fmt"

type ReverseFactorialResult struct {
	finalValue float64
	factorial  int
}

func CalulateReverseFactorial(i int) ReverseFactorialResult {
	x := float64(i)
	y := 1

	for x > 1 {
		x = x / float64((1 + y))
		y += 1
	}

	return ReverseFactorialResult{x, y}
}

func IsFactorial(i int) string {
	res := CalulateReverseFactorial(i)
	if res.finalValue != 1.0 {
		return "NONE"
	} else {
		return fmt.Sprintf("%d!", res.factorial)
	}
}
