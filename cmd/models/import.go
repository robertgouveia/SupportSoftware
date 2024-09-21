package models

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

func Execute(sqlFile string, db *sql.DB, queryName string) {
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

	headers := []string{}
	headers = append(headers, columns...)

	values := make([]interface{}, len(columns))
	for i := range values {
		var value interface{}
		values[i] = &value
	}

	var outputName string
	outputName += strings.ReplaceAll(queryName, " ", "_") + "_"
	outputName += time.Now().Format("2006-01-02-15-04-05")

	file, err := os.Create("./exports/" + outputName + ".csv")
	if err != nil {
		log.Fatalf("Unable to create output file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	err = writer.Write(headers)
	if err != nil {
		log.Fatalf("Error Writing Headers: %v", err)
	}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			log.Fatalf("Error pulling in data: %v", err)
		}

		results := []string{}
		for _, val := range values {
			value := *(val.(*interface{}))
			results = append(results, columnValueToString(value))
		}

		err = writer.Write(results)
		if err != nil {
			log.Printf("Unable to print line: %v, skipping", err)
			continue
		}

	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatalf("Error flushing csv writer: %v", err)
	}

	log.Println("Exported Successfully")
}

func columnValueToString(val interface{}) string {
	switch v := val.(type) {
	case nil:
		return "NULL"
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case string:
		return v
	case []byte:
		return string(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
