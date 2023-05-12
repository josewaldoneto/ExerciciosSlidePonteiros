package main

import "fmt"

// Escreva uma função que receba um ponteiro para um booleano e altere o valor desse
// booleano para o seu inverso.

func invert(ptr *bool) {
	if *ptr == false {
		*ptr = true
	} else {
		*ptr = false
	}
	return
}

func main() {
	b := true
	fmt.Println("Valor booleano é:", b)
	invert(&b)
	fmt.Println("Valor booleano depois da função é:", b)
}
