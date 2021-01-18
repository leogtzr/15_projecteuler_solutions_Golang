package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	is := true

	max := int(math.Sqrt(float64(n)))
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			is = false
			break
		}
	}

	return is
}

func main() {
	sum := 0
	for i := 2; i <= 2_000_000; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
