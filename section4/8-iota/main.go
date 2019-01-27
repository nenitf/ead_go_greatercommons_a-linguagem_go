package main

import "fmt"

// iota são para constantes int cujo seguirão uma mesma regra
// caso seja posto somente iota será adicionado +1
// também é possivel criar sua propria regra, como aumentar em 2
// blank indentifiers são contabilizados

const (
	a = iota + 2
	_
	c
	d
	e
	_
	g
)

func main() {
	fmt.Println(a, c, d, e, g)
}
