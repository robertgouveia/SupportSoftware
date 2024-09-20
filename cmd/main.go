package main

import (
	"database/cmd/models"
	"fmt"
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

	conn, db := models.Connection{}.Open()
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	chosenDB := models.Question(conn.Databases(db), "You chose", true, true)
	output, err := db.Exec("USE " + chosenDB)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	log.Println(output)

	imports := models.ImportList{}.Get()
	queryName := models.Question(models.ImportNames(imports), "You chose", false, false)

	var sqlFile string
	for _, i := range imports {
		if queryName == models.ImportToName(i.Name) {
			sqlFile = i.Path
		}
	}

	if sqlFile == "" {
		log.Fatal("Unable to find SQL File")
	}

	queries, err := os.ReadFile("./imports/" + sqlFile)
	if err != nil {
		log.Fatalf("Unable to read SQL: %v", err)
	}

	rows, err := db.Query(string(queries))
	if err != nil {
		log.Fatalf("Unable to execute query: %v", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("Unable to count rows: %v", err)
	}

	values := make([]interface{}, len(columns))
	for i := range values {
		var value interface{}
		values[i] = &value
	}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			log.Fatalf("Error pulling in data: %v", err)
		}

		for i, col := range values {
			fmt.Printf("%s: %v\n", columns[i], *(col.(*interface{})))
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}

}
