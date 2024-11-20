package client

import (
	"api/pkg/models"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func StartClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler body da requisição")
		panic(err)
	}

	os.Stdout.Write(body)

	// salvar em arquivo
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var change models.UsdBrl
	err = json.Unmarshal(body, &change)
	if err != nil {
		log.Println("Erro ao decodificar JSON")
		panic(err)
	}

	changeData, err := json.MarshalIndent(change, "", " ")
	if err != nil {
		log.Println("Erro ao formatar resposta da api")
		panic(err)
	}

	_, err = f.WriteString(string(changeData))
	if err != nil {
		log.Println("Erro ao salvar arquivo")
		panic(err)
	}

	log.Println("Cotação salva no .txt")

}
