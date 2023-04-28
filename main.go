package main

import (
	"log"

	"github.com/Marian86029/Bali2024/db"
	"github.com/Marian86029/Bali2024/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin Conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
