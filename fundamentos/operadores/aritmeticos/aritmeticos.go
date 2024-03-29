package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("soma =>", a+b)
	fmt.Println("subtração =>", a-b)
	fmt.Println("divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("OU =>", a|b)  // 11 | 10 = 11
	fmt.Println("XOR =>", a^b) // 11 ^ 10 = 01

	c := 4.0
	d := 2.0

	//outras operacoes usando math...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))

}
