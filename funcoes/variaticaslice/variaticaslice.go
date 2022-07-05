package main

import "fmt"

// imprimindo uma função variática com slice
func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...)
}

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}
