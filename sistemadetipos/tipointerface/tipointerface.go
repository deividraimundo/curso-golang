package main

import "fmt"

type curso struct {
	nome string
}

type imprimivel interface {
	toString() string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}
	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)

	var coisa3 imprimivel = curso{"Golang"}
	imprimir(coisa3)
}

func (c curso) toString() string {
	return c.nome
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}
