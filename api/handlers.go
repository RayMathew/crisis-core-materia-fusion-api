package main

import (
	"errors"
	"fmt"
	"net/http"

	ccmf "github.com/RayMathew/crisis-core-materia-fusion-api/internal/crisis-core-materia-fusion"
	"github.com/RayMathew/crisis-core-materia-fusion-api/internal/request"
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
	var allDisplayMateria []ccmf.MateriaDTO
	var allMateria []ccmf.Materia
	var err error

	allMateria, err = app.getAllMateriaFromApprSource()

	if err != nil {
		app.serverError(w, r, err)
	}

	for _, materia := range allMateria {
		allDisplayMateria = append(allDisplayMateria, ccmf.MateriaDTO{
			Name:        materia.Name,
			Type:        materia.DisplayType,
			Description: materia.Description,
		})
	}

	err = response.JSON(w, http.StatusOK, allDisplayMateria)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) fuseMateria(w http.ResponseWriter, r *http.Request) {
	var fusionReq ccmf.MateriaFusionRequest
	err := request.DecodeJSON(w, r, &fusionReq)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	var allMateria []ccmf.Materia

	allMateria, err = app.getAllMateriaFromApprSource()

	if err != nil {
		app.serverError(w, r, err)
	}

	var materia1Type string
	var materia1Grade int
	var materia2Type string
	var materia2Grade int

	// app.logger.Info(fusionReq.Materia1Name)
	// app.logger.Info(fusionReq.Materia2Name)

	for _, materia := range allMateria {
		// app.logger.Info(materia.Name)
		// app.logger.Info(materia.Type)
		if materia1Type != "" && materia2Type != "" {
			app.logger.Info("break")
			break
		}
		if materia.Name == fusionReq.Materia1Name && materia1Type == "" {
			materia1Type = materia.Type
			materia1Grade = materia.Grade
		}
		if materia.Name == fusionReq.Materia2Name && materia2Type == "" {
			materia2Type = materia.Type
			materia2Grade = materia.Grade
		}
	}

	// app.logger.Info(materia1Type)
	// app.logger.Info(materia2Type)

	if materia1Type == "" || materia2Type == "" {
		app.badRequest(w, r, errors.New("one or both of the Materia names not recognised"))
		return
	}

	exchangePositionsIfNeeded(&fusionReq, &materia1Grade, &materia2Grade, &materia1Type, &materia2Type)
	// app.logger.Info("positions after shuffle")
	// app.logger.Info("materia 1")
	// fmt.Println(materia1Grade, materia1Type, fusionReq.Materia1Name, fusionReq.Materia1Mastered)
	// app.logger.Info("materia 2")
	// fmt.Println(materia2Grade, materia2Type, fusionReq.Materia2Name, fusionReq.Materia2Mastered)

	relevantBasicRuleMap := ccmf.BasicRuleMap[ccmf.MateriaType(materia1Type)]
	var relevantBasicRule ccmf.BasicCombinationRule

	for _, rule := range relevantBasicRuleMap {
		if rule.SecondMateriaType == ccmf.MateriaType(materia2Type) {
			relevantBasicRule = rule
			break
		}
	}

	var resultantMateria ccmf.MateriaDTO
	resultantMateriaGrade := determineGrade(fusionReq, materia1Grade)

	if relevantBasicRule.FirstMateriaType == "" {
		app.logger.Info("none of the basic rules satisfy the requirement.")

		//get final output using complex rules
		resultantMateria = useComplexRules(fusionReq, materia1Grade, materia2Grade, materia1Type, materia2Type)
	} else {
		//get final output using basic rules
		fmt.Println(materia1Grade, materia2Grade, materia1Type, materia2Type)
		resultantMateriaType := relevantBasicRule.ResultantMateriaType
		fmt.Println(resultantMateriaGrade, resultantMateriaType)

		for _, materia := range allMateria {
			if materia.Grade == resultantMateriaGrade && materia.Type == string(resultantMateriaType) {
				resultantMateria.Name = materia.Name
				resultantMateria.Type = materia.DisplayType
				resultantMateria.Description = materia.Description
				break
			}
		}

	}

	err = response.JSON(w, http.StatusOK, resultantMateria)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// positions are exchange if materia2 grade is higher
func exchangePositionsIfNeeded(fusionReq *ccmf.MateriaFusionRequest, materia1Grade, materia2Grade *int, materia1Type, materia2Type *string) {
	if *materia2Grade > *materia1Grade {
		tempGrade := *materia1Grade
		*materia1Grade = *materia2Grade
		*materia2Grade = tempGrade

		tempType := *materia1Type
		*materia1Type = *materia2Type
		*materia2Type = tempType

		tempName := fusionReq.Materia1Name
		fusionReq.Materia1Name = fusionReq.Materia2Name
		fusionReq.Materia2Name = tempName

		tempMastered := fusionReq.Materia1Mastered
		fusionReq.Materia1Mastered = fusionReq.Materia2Mastered
		fusionReq.Materia2Mastered = tempMastered
	}
}

func determineGrade(req ccmf.MateriaFusionRequest, materia1Grade int) int {
	finalGrade := materia1Grade

	if finalGrade != 8 && req.Materia1Mastered {
		finalGrade += 1
	}
	if finalGrade != 8 && req.Materia2Mastered {
		finalGrade += 1
	}
	return finalGrade
}

func useComplexRules(fusionReq ccmf.MateriaFusionRequest, materia1Grade, materia2Grade int, materia1Type, materia2Type string) (resultantMateria ccmf.MateriaDTO) {

	return
}

func (app *application) getAllMateriaFromApprSource() (allMateria []ccmf.Materia, err error) {

	// Check if allMateria data is in cache
	if data, found := app.getCachedData(string(ccmf.AllMateriaCacheKey)); found {
		// Type assertion: assert that data is of type []Materia
		if allMateriaCache, ok := data.([]ccmf.Materia); ok {
			allMateria = allMateriaCache
			app.logger.Debug("cache hit")
		} else {
			app.logger.Error("Failed to assert cached data as []Materia")
			return nil, errors.New("this is a simple error message")
		}
	} else {
		// allMateria data is not in cache. Get from DB
		app.logger.Debug("cache miss")
		allMateria, err = app.db.GetAllMateria()
		app.setCache(string(ccmf.AllMateriaCacheKey), allMateria)
	}
	return
}
