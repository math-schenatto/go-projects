package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Matheus Schenatto"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func strpoint() {
	matheus := Cliente{nome: "Matheus"}
	matheus.andou()

	fmt.Printf("O valor da struct com nome é %v", matheus.nome)
}
