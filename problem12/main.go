package main

import "fmt"

func main() {
	n := 1
	triangle := 0
	for {
		triangle += n
		numberOfDivisors := 0
		for i := 1; i <= triangle; i++ {
			if triangle%i == 0 {
				numberOfDivisors++
			}
		}
		if numberOfDivisors >= 500 {
			fmt.Printf("triangle: %d has %d divisors\n", triangle, numberOfDivisors)
			break
		}
		n++
	}
}
