package models

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func ExecuteSQL(sqlFile string, db *sql.DB, queryName string) {
	file, err := os.Open("./imports/" + sqlFile)
	if err != nil {
		log.Fatalf("Unable to read SQL File: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Fatalf("Unable to read SQL File: %v", err)
	}

	queryFile, err := os.Create("./imports/temp.sql")
	if err != nil {
		log.Fatalf("Failure creating temp file: %v", err)
	}
	defer queryFile.Close()
	writer := bufio.NewWriter(queryFile)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "DECLARE") {
			parts := strings.Split(line, "=")
			name := strings.Split(strings.Split(parts[0], "@")[1], " ")[0]
			if len(parts) > 1 {
				value := strings.ReplaceAll(strings.TrimSpace(parts[1]), ";", "")
				p := tea.NewProgram(TextInputModel("Value", name))
				if _, err := p.Run(); err != nil {
					log.Fatal(err)
				}
				_, err := writer.WriteString(strings.ReplaceAll(line, value, TextInputModel("", "").GetInput()) + "\n")
				if err != nil {
					log.Fatalf("Unable to write to temp file: %v", err)
				}
			}
			continue
		}
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Unable to write to temp file: %v", err)
		}
	}

	err = writer.Flush()
	if err != nil {
		log.Fatalf("Unable to flush the temp file: %v", err)
	}

	Execute("temp.sql", db, queryName)

	err = os.Remove("./imports/temp.sql")
	if err != nil {
		log.Fatalf("Unable to remove temp file: %v", err)
	}
}
