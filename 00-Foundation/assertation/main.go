package main

import "fmt"

func main() {
	var x interface{} = "Matheus Schenatto"
	println(x.(string))

	res, ok := x.(int)

	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}
