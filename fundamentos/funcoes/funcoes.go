package main

import "fmt"

func main() {
	resultado := somar(3, 4)
	imprimir(int(resultado))
}

func somar(number1 float64, number2 float64) float64 {
	return number1 + number2
}

func imprimir(value int) {
	fmt.Println(value)
}
