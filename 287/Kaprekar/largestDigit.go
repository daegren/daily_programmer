package Kaprekar

import "math"

func LargestDigit(i int) int {
	largestDigit := 0
	floatInput := float64(i)
	for exp := 3; exp >= 0; exp-- {
		val := int(floatInput / math.Pow10(exp))
		if val > largestDigit {
			largestDigit = val
		}
		floatInput = floatInput - (float64(val) * math.Pow10(exp))
	}

	return largestDigit
}
