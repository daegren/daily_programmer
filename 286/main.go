package main

import (
	"fmt"

	"github.com/daegren/daily_programmer/286/ReverseFactorial"
)

func main() {
	input := [...]int{3628800, 479001600, 6, 18}
	for _, element := range input {
		fmt.Printf("Reverse Factorial of %d is %s\n", element, ReverseFactorial.IsFactorial(element))
	}
}
