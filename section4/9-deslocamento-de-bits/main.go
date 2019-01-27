package main

import "fmt"

func main() {
	x := 4

	// desloca duas casas para <-
	// "aumentam 00"
	y := x << 2

	// desloca duas casas para ->
	// "diminuem 00"
	z := x >> 2
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
}
