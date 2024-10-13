package main

import "fmt"

func main() {
	salarios := map[string]int{"Matheus": 100, "Wesley": 200}
	// fmt.Println(salarios["Matheus"])
	// delete(salarios, "Wesley")
	// salarios["Wes"] = 5000
	// fmt.Println(salarios["Wes"])

	// sal := make(map[string]int)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d \n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d \n", salario)
	}
}
