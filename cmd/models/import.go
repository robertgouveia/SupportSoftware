package models

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Import struct {
	Name string
	Path string
}

func ImportNames(i []Import) []string {
	names := []string{}

	for _, name := range i {
		names = append(names, name.Name)
	}

	return names
}

func ImportToName(file string) string {
	f := filepath.Base(file)
	f = strings.TrimSuffix(f, ".sql")
	f = strings.ReplaceAll(f, "_", " ")
	return f
}

func Execute(sqlFile string, db *sql.DB) {
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
	}
}
