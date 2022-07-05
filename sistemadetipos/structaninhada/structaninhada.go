package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{
				produtoID:  1,
				quantidade: 2,
				preco:      12.10,
			},
			{
				produtoID:  2,
				quantidade: 1,
				preco:      23.49,
			},
			{
				produtoID:  11,
				quantidade: 100,
				preco:      3.13,
			},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f\n", pedido.valorTotal())
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total
}
