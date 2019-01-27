package main

import "fmt"

func main() {
	s := "ascii éøâ 香"

	// In Go, a string is in effect a read-only slice of bytes
	// https://blog.golang.org/strings

	// caractere por caractere
	for _, v := range s {
		fmt.Printf("%b - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	// byte por byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}
