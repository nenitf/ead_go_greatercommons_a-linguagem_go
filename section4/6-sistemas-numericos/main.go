package main

import "fmt"

func main() {
	// https://godoc.org/fmt
	a := 106
	fmt.Printf("valor\tdecimal[10]\tbinario[2]\thexadecimal[16]\n")
	fmt.Printf("%v\t%d\t%b\t%#x\n", a, a, a, a)

}
