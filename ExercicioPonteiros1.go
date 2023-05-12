package main

import "fmt"

// Escreva uma função swap que receba dois ponteiros para int como argumentos e troque
// os valores apontados por eles.

func swap(ptrx *int, ptry *int) {
	valorX := *ptrx
	valorY := *ptry
	*ptrx = valorY
	*ptry = valorX
}
func main() {
	x := 10
	y := 99
	fmt.Println("Valor de X antes da troca", x)
	fmt.Println("Valor de Y antes da troca", y)
	swap(&x, &y)
	fmt.Println("Valor de X depois da troca", x)
	fmt.Println("Valor de Y depois da troca", y)
}
