package models

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ImportList struct {
	Imports []Import
}

func (i ImportList) Get() []Import {
	i.Imports = []Import{}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working Directory: %v", err)
	}

	relativeDir := filepath.Join(wd, "/imports")

	files, err := os.ReadDir(relativeDir)
	if err != nil {
		log.Fatalf("Error reading import file directory: %v", err)
	}

	for _, file := range files {
		f := filepath.Base(file.Name())
		f = strings.TrimSuffix(f, ".sql")
		f = strings.ReplaceAll(f, "_", " ")

		i.Imports = append(i.Imports, Import{Name: f, Path: file.Name()})
	}

	return i.Imports
}
