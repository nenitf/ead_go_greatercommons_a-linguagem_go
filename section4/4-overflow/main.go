package main

import "fmt"

func main() {
	// https://golang.org/pkg/builtin/#uint16
	// uint16 is the set of all unsigned 16-bit integers
	// Range: 0 through 65535
	var i uint16
	// i = 65536 não é possivel
	i = 65535 // máximo
	fmt.Println(i)
	i++ // 65536? não! 0
	fmt.Println(i)
	i++ // 1
	fmt.Println(i)
	i++ // 2
	fmt.Println(i)
}
