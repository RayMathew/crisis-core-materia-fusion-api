package main

import (
	"net/http"

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
	user, err := app.db.GetAllMateria()

	if err != nil {
		app.serverError(w, r, err)
	}

	err = response.JSON(w, http.StatusOK, user)
	if err != nil {
		app.serverError(w, r, err)
	}
}
