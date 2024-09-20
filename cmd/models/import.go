package models

import (
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
