package main

import "fmt"

func main() {
	n := 1

	inc1(n) // passagem por valor
	fmt.Println(n)

	// revisão: & usado para obter o endereço da variavel
	inc2(&n)
	fmt.Println(n)
}

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *

func inc2(n *int) {
	// revisar: * é usado para desreferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta
	*n++
}
