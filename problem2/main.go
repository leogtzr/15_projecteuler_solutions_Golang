package main

import "fmt"

func main() {
	a := 0
	b := 1
	sum := 0
	const max = 4_000_000

	for b < max {
		buff := a + b
		if b%2 == 0 {
			sum += b
		}
		a = b
		b = buff
	}

	fmt.Printf("The sum is: [%d]\n", sum)
}
