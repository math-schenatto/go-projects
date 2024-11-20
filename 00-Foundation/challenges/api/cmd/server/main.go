package main

import (
	"api/internal/db"
	"api/internal/server"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o servidor...")
	db.InitDB()
	server.StartServer()
}
