package main

import "fmt"

func isDivisibleBy3Or5(n int) bool {
	return n%3 == 0 || n%5 == 0
}

func main() {
	sum := 0
	for i := 1; i < 1_000; i++ {
		if isDivisibleBy3Or5(i) {
			sum += i
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}
