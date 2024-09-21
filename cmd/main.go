package main

import (
	"bufio"
	"database/cmd/models"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	file, err := os.Open("./connection/connection.txt")
	if err != nil {
		log.Fatalf("Unable to get connection list: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	connections := []string{}
	for scanner.Scan() {
		connection := scanner.Text()
		connections = append(connections, os.Getenv("SERVER")+"."+connection)
	}

	os.Setenv("SERVER", models.Question(connections, "Select a Server", "You chose", true, true))

	conn, db := models.Connection{}.Open()
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	os.Setenv("DATABASE", models.Question(conn.Databases(db), "Select a Database", "You chose", true, true))
	conn, db = models.Connection{}.Open()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	imports := models.ImportList{}.Get()
	queryName := models.Question(models.ImportNames(imports), "Select an Export", "You chose", false, false)

	var sqlFile string
	for _, i := range imports {
		if queryName == models.ImportToName(i.Name) {
			sqlFile = i.Path
		}
	}

	if sqlFile == "" {
		log.Fatal("No SQL File Selected")
	}

	models.Execute(sqlFile, db, queryName)
}
