package main

import (
	"fmt"
	"math"
)

func largest(factors []int) int {
	max := factors[0]

	for _, factor := range factors {
		if factor > max {
			max = factor
		}
	}

	return max
}

func primeFactors(n int) []int {
	factors := make([]int, 0)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
			i--
		}
	}
	if n > 0 {
		factors = append(factors, n)
	}

	return factors
}

func main() {
	n := 600851475143
	largestPrimeFactor := largest(primeFactors(n))
	fmt.Println(largestPrimeFactor)
}
