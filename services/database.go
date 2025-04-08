package services

import (
	"database/sql"
	"fmt"
	"log"
)

// database handler
var db *sql.DB

func InitDB() {
	// capture connection properties
	dsn := "postgres://postgres:postgres@localhost/hubcook?sslmode=disable"

	// connection a la database
	db, err := sql.Open("postgres", dsn)

	// tentative de connection a la db
	if err != nil {
		log.Fatalf("Erreur lors de la conection a la DB : %v", err)
	}

	fmt.Println("Connexion database OK")
}
