package main

// MateriaDTO provides Materia details - Name, Description and Type (Magic / Command / Support / Independent)
type MateriaDTO struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
