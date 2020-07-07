package main

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p_produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

}
