package main

import (
	"github.com/RayMathew/crisis-core-materia-fusion-api/internal/validator"
)

type MateriaFusionRequest struct {
	Materia1Mastered *bool               `json:"materia1mastered"`
	Materia2Mastered *bool               `json:"materia2mastered"`
	Materia1Name     string              `json:"materia1name"`
	Materia2Name     string              `json:"materia2name"`
	Validator        validator.Validator `json:"-"`
}
