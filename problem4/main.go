package main

import "fmt"

func isPalindrome(n int) bool {
	reverse := 0
	original := n
	for n != 0 {
		reverse = (reverse * 10) + (n % 10)
		n /= 10
	}
	return original == reverse
}

func main() {
	max := 999
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			if isPalindrome(product) && product > max {
				max = product
			}
		}
	}

	fmt.Println(max)
}
