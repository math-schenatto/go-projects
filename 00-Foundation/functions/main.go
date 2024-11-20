package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}

func sum2(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {
	valor, err := sum(50, 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("A soma é: ", valor)

	fmt.Println(sum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
