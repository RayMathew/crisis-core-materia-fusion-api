package main

// MateriaDTO provides Materia details - Name, Description and Type (Magic / Command / Support / Independent)
type MateriaDTO struct {
	Name        string `json:"name" example:"Thunder"`
	Type        string `json:"type" example:"Magic"`
	Description string `json:"description" example:"Shoots lightning forward dealing thunder damage."`
}

// StatusDTO provides status of the server
type StatusDTO struct {
	Status string `json:"Status" example:"OK"`
}

// ErrorResponseDTO provides Error message
type ErrorResponseDTO struct {
	Error string `json:"Error" example:"The server encountered a problem and could not process your request"`
}
