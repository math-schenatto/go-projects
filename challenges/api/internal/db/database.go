package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./requests.db")
	if err != nil {
		log.Fatal(err)
	}

	// Tentativa de conexão com o banco de dados
	for retries := 0; retries < 5; retries++ {
		if err = DB.Ping(); err == nil {
			break
		}
		log.Println("Tentando conectar ao banco de dados...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Não foi possível conectar ao banco de dados: ", err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS usdbrl (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        code TEXT,
        codein TEXT,
        name TEXT,
        high TEXT,
        low TEXT,
        varBid TEXT,
        pctChange TEXT,
        bid TEXT,
        ask TEXT,
        timestamp TEXT,
        createDate TEXT
    );
    `

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database conection success...")

}
