package main

import "fmt"

// Escreva uma função que receba um ponteiro para uma struct que contém informações
// de um produto (nome, preço e quantidade em estoque).
// A função deve atualizar o preço
// desse produto para um novo valor fornecido como argumento.

type Produto struct {
	nome       string
	preco      float64
	qntEstoque int
}

func novoPreco(p *Produto, newPrice float64) {
	p.preco = newPrice
}

func main() {
	produto := Produto{nome: "Leite", preco: 6.99, qntEstoque: 100}
	fmt.Println("Preço antigo:", produto.preco)
	novoPreco(&produto, 8.99)
	fmt.Println("Preço novo:", produto.preco)
}
