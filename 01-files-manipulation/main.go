package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// ESCREVENDO STRING
	// tamanho, err := f.WriteString("Hello World!")
	// if err != nil {
	// 	panic(err)
	// }

	// ESCREVENDO BYTES
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes", tamanho)
	f.Close()

	/// LEITURA
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// arquivo recebe bytes

	fmt.Println(string(arquivo))

	// arquivo deve ser lido em chunks
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read((buffer))
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	arquivo2.Close()

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
