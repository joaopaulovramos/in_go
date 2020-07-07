package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // Retorno Limpo
}

func main() {

	fmt.Println(trocar(4, 5))
}
