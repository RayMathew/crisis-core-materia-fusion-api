package crisiscoremateriafusion

import (
	"fmt"
	"strings"
)

type MateriaFusionRequest struct {
	Materia1Name     string `json:"materia1name"`
	Materia1Mastered *bool  `json:"materia1mastered"`
	Materia2Name     string `json:"materia2name"`
	Materia2Mastered *bool  `json:"materia2mastered"`
}

func validateNameRequired(field, fieldName string) error {
	if strings.TrimSpace(field) == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

func validateMasteredRequired(field *bool, fieldName string) error {
	if field == nil {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

func (req *MateriaFusionRequest) ValidateUserRequest() error {
	if err := validateNameRequired(req.Materia1Name, "materia1name"); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	if err := validateNameRequired(req.Materia2Name, "materia2name"); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	if err := validateMasteredRequired(req.Materia1Mastered, "materia1mastered"); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	if err := validateMasteredRequired(req.Materia2Mastered, "materia2mastered"); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}
