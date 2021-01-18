package main

import "fmt"

func main() {
	a := 0
	b := 1

	for i := 0; i < 10; i++ {
		ret := a
		a, b = b, a+ret
		fmt.Println(ret)
	}
}
