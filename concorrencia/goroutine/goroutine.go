package main

import (
	"fmt"
	"time"
)

func main() {
	// fale("Maria", "Por que você não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)
	// time.Sleep(time.Second * 5)

	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)
}

func fale(pessoa, texto string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}
