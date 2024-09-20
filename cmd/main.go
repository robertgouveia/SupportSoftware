package main

import (
	"database/cmd/models"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	conn, db := models.Connection{}.Open()
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	choice := models.Question(conn.Databases(db), "You chose")

	log.Println(choice)
}
