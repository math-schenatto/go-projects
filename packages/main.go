package main

import (
	"fmt"

	"github.com/mathesche/curso-go/matematica"
)

// go build
// GOOS=windows go build main.go
// go env
func main() {
	soma := matematica.Soma(10, 20)

	fmt.Println("Resultado: ", soma)

	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Andar())
}
