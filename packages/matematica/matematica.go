package matematica

import "fmt"

// Só fica disponivel nos outros pacotes se for a inicial maiucula
func Soma[T int | float64](a T, b T) T {
	return a + b
}

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	fmt.Println(c.Marca)
	return "Carro andando"
}

// go get nome_do_pacote
// go get tidy - organiza o modulo, remove pacotes não utilizados e baixa os não baixados
