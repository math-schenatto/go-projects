package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func (c Cliente) andou() {
	c.nome = "Matheus Schenatto"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	matheus := Cliente{nome: "Matheus"}
	matheus.andou()

	conta := Conta{saldo: 100}

	conta.simular(200)

	fmt.Printf("O valor da struct com nome é %v\n", matheus.nome)
	fmt.Printf("O valor do saldo é %d", conta.saldo)
}
