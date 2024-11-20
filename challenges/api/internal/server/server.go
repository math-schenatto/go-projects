package server

import (
	"api/internal/db"
	"api/pkg/models"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func StartServer() {
	http.HandleFunc("/cotacao", ChangeOfDayHandler)
	// Log indicando que o servidor está rodando
	log.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}

func ChangeOfDayHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	change, error := ChangeOfDay(ctx)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	dbCtx, dbCancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer dbCancel()

	// Tentando persistir os dados no banco
	select {
	case <-dbCtx.Done():
		log.Println("Database timeout")
		http.Error(w, "Timeout ao salvar dados no banco", http.StatusRequestTimeout)
		return
	default:
		// Persistindo os dados
		error = persistUsdBrl(dbCtx, *change)
		if error != nil {
			log.Println("Erro ao salvar no banco de dados:", error)
			http.Error(w, "Erro ao salvar dados no banco", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(change); err != nil {
		log.Println("Erro ao codificar a resposta:", err)
		http.Error(w, "Erro ao enviar resposta", http.StatusInternalServerError)
	}
}

func ChangeOfDay(ctx context.Context) (*models.UsdBrl, error) {
	log.Println("Iniciando Requisição")
	req, error := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if error != nil {
		return nil, error
	}

	// Faz a requisição e captura a resposta
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Erro na requisição ", err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Println("Requisição completa")
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var usdBrl models.UsdBrl
	error = json.Unmarshal(body, &usdBrl)
	if error != nil {
		return nil, error
	}

	return &usdBrl, nil
}

func persistUsdBrl(ctx context.Context, data models.UsdBrl) error {

	query := `
		INSERT INTO usdbrl (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, createDate)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	// Para executar query com contexto
	stmt, err := db.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Erro ao preparar statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		data.USDBRL.Code,
		data.USDBRL.Codein,
		data.USDBRL.Name,
		data.USDBRL.High,
		data.USDBRL.Low,
		data.USDBRL.VarBid,
		data.USDBRL.PctChange,
		data.USDBRL.Bid,
		data.USDBRL.Ask,
		data.USDBRL.Timestamp,
		data.USDBRL.CreateDate,
	)
	if err != nil {
		log.Printf("Erro ao salvar a requisição: %v", err)
		return err
	}

	log.Println("Requisição salva com sucesso!")
	return nil
}
