package main

import (
	"database/cmd/models"
	"log"

	"github.com/charmbracelet/bubbles/list"
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

	items := []list.Item{}
	databases := conn.Databases(db)
	for _, db := range databases {
		items = append(items, db)
	}

	models.DBList{}.New(items, "Select a Database: ")
}
