package models

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
