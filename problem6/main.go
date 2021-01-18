package main

import "fmt"

func main() {
	sumSquares := 0
	squaresOfTheSum := 0
	for i := 1; i <= 100; i++ {
		sumSquares += i * i
		squaresOfTheSum += i
	}
	diffSums := (squaresOfTheSum * squaresOfTheSum) - sumSquares
	fmt.Println(diffSums)
}
