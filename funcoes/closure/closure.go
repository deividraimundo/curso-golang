package main

import "fmt"

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}
