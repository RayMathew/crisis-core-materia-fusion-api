package crisiscoremateriafusion

type MateriaFusionRequest struct {
	Materia1Name     string `json:"materia1name"`
	Materia1Mastered bool   `json:"materia1mastered"`
	Materia2Name     string `json:"materia2name"`
	Materia2Mastered bool   `json:"materia2mastered"`
}
