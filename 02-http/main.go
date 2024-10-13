package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get()
	if err != nil {
		panic(err)
	}

	//defer atrasa a execução do código na linha, executado no final do programa
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))

}
