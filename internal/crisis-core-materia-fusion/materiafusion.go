package crisiscoremateriafusion

type Materia struct {
	Name        string `json:"name"`
	Type        string `json:"materia_type"`
	Grade       int    `json:"grade"`
	DisplayType string `json:"display_type"`
	Description string `json:"description"`
}

type Rule struct {
	FirstType     string `json:"first_type"`
	SecondType    string `json:"second_type"`
	ResultantType string `json:"resultant_type"`
}

type MateriaFusionService interface {
	GetAllMateria() ([]Materia, error)
	GetAllRules() ([]Rule, error)
}
