package main

/*
   Crie um tipo. O tipo subjacente deve ser int.
   Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
   Na função main:

   Demonstre o valor da variável "x"
   Demonstre o tipo da variável "x"
   Atribua 42 à variável "x" utilizando o operador "="
   Demonstre o valor da variável "x"

   Para os aventureiros: https://golang.org/ref/spec#Types
*/

import (
	"fmt"
	"reflect"
)

type unt int

var x unt

func main() {
	fmt.Printf("%v\t%v\t%T\n", x, reflect.TypeOf(x), x)
	x = 42
	fmt.Println(x)
}
