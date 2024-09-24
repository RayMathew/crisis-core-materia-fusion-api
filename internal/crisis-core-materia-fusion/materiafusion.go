package crisiscoremateriafusion

type Materia struct {
	Name        string `json:"name"`
	Type        string `json:"materia_type"`
	Grade       int    `json:"grade"`
	DisplayType string `json:"display_type"`
	Description string `json:"description"`
}

type UserService interface {
	GetAllMateria() ([]Materia, error)
}
