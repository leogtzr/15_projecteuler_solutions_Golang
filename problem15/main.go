// func factorial(n int) uint64 {
//  var f uint64 = 1

//  for i := 1; i <= n; i++ {
//      f *= uint64(i)
//  }

//  return f
// }

package main

import (
	"fmt"
	"math/big"
)

func factorial(n int) *big.Int {
	f := big.NewInt(1)

	for i := 1; i <= n; i++ {
		f.Mul(f, big.NewInt(int64(i)))
	}

	return f
}

func main() {
	// n = 20; (2 * n)! / (n!)^2

	n := 20
	a := factorial(2 * n)
	b := factorial(n)
	b.Mul(b, b)

	result := new(big.Int)

	result.Div(a, b)

	fmt.Println(result)
}
