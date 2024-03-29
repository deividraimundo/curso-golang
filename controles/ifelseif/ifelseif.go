package main

import "fmt"

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.1))
	fmt.Println(notaParaConceito(3))
	fmt.Println(notaParaConceito(11))
}

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else if n >= 0 && n < 3 {
		return "E"
	} else {
		return "Erro - Digite um número de 0 a 10"
	}
}
