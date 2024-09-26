package main

import (
	"net/http"

	crisiscoremateriafusion "github.com/RayMathew/crisis-core-materia-fusion-api/internal/crisis-core-materia-fusion"
	"github.com/RayMathew/crisis-core-materia-fusion-api/internal/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getAllMateria(w http.ResponseWriter, r *http.Request) {
	allMateria, err := app.db.GetAllMateria()

	var allDisplayMateria []crisiscoremateriafusion.MateriaDTO

	for _, materia := range allMateria {
		allDisplayMateria = append(allDisplayMateria, crisiscoremateriafusion.MateriaDTO{
			Name:        materia.Name,
			Type:        materia.DisplayType,
			Description: materia.Description,
		})
	}

	if err != nil {
		app.serverError(w, r, err)
	}

	err = response.JSON(w, http.StatusOK, allDisplayMateria)
	if err != nil {
		app.serverError(w, r, err)
	}
}
