package main

import (
	"fmt"
)

func main() {
	z := kara(1234, 5678)
	fmt.Println("kara:", z)
	fmt.Println("digits: ", numDigits(475912))
}

func kara(x, y int) int {
	numDigs := numDigits(x)
	return x * y
}

func numDigits(number int) int {
	digits := 0
	for number >= 1 {
		number = number / 10
		digits++
	}
	return digits
}
