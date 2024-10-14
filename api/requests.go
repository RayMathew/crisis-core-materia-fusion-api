package main

import (
	"github.com/RayMathew/crisis-core-materia-fusion-api/internal/validator"
)

// MateriaFusionRequest provides input Materia names and their Mastered states
type MateriaFusionRequest struct {
	Materia1Mastered *bool               `json:"materia1mastered" example:"true"`
	Materia2Mastered *bool               `json:"materia2mastered" example:"false"`
	Materia1Name     string              `json:"materia1name" example:"Fire"`
	Materia2Name     string              `json:"materia2name" example:"Blizzard"`
	Validator        validator.Validator `json:"-"`
}
