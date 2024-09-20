package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

type Connection struct {
	Server   string
	User     string
	Password string
	Port     string
	Database string
}

func (c *Connection) New() (*Connection, error) {
	c.Server = os.Getenv("SERVER")
	c.User = os.Getenv("USERNAME")
	c.Password = os.Getenv("PASSWORD")
	c.Port = os.Getenv("PORT")
	c.Database = os.Getenv("DATABASE")

	if c.Server == "" || c.User == "" || c.Password == "" || c.Port == "" || c.Database == "" {
		return nil, errors.New("Unable to retrieve database values.")
	}

	return c, nil
}

func (c *Connection) ConnectionString() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", c.User, c.Password, c.Server, c.Port, c.Database)
}

func (c Connection) Open() (*Connection, *sql.DB) {
	conn, err := c.New()
	if err != nil {
		log.Panic(err)
	}

	db, err := sql.Open("sqlserver", conn.ConnectionString())
	if err != nil {
		log.Fatalf("Error creating connection pool: %v", err)
	}

	return conn, db
}

func (c *Connection) Databases(db *sql.DB) []Option {
	query := "SELECT name FROM sys.databases"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Query Failed: %v", err)
	}
	defer rows.Close()

	databases := []Option{}

	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			log.Fatalf("Unable to scan row: %v", err)
		}
		databases = append(databases, Option(dbName))
	}

	return databases
}
