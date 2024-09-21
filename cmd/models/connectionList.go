package models

import (
	"bufio"
	"log"
	"os"
)

func GetConnections() []string {
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

	return connections
}
