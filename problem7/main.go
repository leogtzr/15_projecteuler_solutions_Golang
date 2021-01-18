package main

import (
	"fmt"
	"math"
)

func nThPrimeNumber(n int) int {
	if n == 0 {
		return 0
	}
	count := 0
	i := 2
	for ; ; i++ {
		if isPrime(i) {
			count++
		}
		if count == n {
			break
		}
	}

	return i
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	is := true

	max := math.Sqrt(float64(n))
	for i := 2; i <= int(max); i++ {
		if n%i == 0 {
			is = false
			break
		}
	}

	return is
}

func main() {
	fmt.Println(nThPrimeNumber(100_001))
}
