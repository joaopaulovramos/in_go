package main

import "fmt"

func clousure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)
	imprimrX := clousure()
	imprimrX()
}
