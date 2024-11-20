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

type Pessoa interface {
	Desativar()
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	matheus := Cliente{
		Nome:  "Matheus",
		Idade: 28,
		Ativo: true,
	}

	Desativacao(matheus)

	fmt.Println(matheus.Nome, matheus.Idade, matheus.Ativo)
}
