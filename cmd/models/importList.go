package models

import (
	"log"
	"os"
	"path/filepath"
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
		f := ImportToName(file.Name())
		i.Imports = append(i.Imports, Import{Name: f, Path: file.Name()})
	}

	return i.Imports
}
