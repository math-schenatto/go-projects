package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     string
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // composição de struct
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {
	matheus := Cliente{
		Nome:  "Matheus",
		Idade: 28,
		Ativo: true,
	}

	matheus.Desativar()

	fmt.Println(matheus.Nome, matheus.Idade, matheus.Ativo)
}
