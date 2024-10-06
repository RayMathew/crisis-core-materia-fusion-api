package main

import (
	"errors"
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

	err = fusionReq.ValidateUserRequest()
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

	app.logger.Debug("start")

	for _, materia := range allMateria {
		if materia1Type != "" && materia2Type != "" {
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

	if materia1Type == "" || materia2Type == "" {
		app.badRequest(w, r, errors.New("one or both of the Materia names are not recognised"))
		return
	}

	exchangePositionsIfNeeded(&fusionReq, &materia1Grade, &materia2Grade, &materia1Type, &materia2Type)

	relevantBasicRuleMap := ccmf.BasicRuleMap[ccmf.MateriaType(materia1Type)]
	var relevantBasicRule ccmf.BasicCombinationRule

	for _, rule := range relevantBasicRuleMap {
		if (rule.FirstMateriaType == ccmf.MateriaType(materia1Type)) &&
			(rule.SecondMateriaType == ccmf.MateriaType(materia2Type)) {
			relevantBasicRule = rule
			break
		}
	}

	var resultantMateria ccmf.MateriaDTO
	resultantMateriaGrade := determineGrade(fusionReq, materia1Grade)

	if relevantBasicRule.FirstMateriaType == "" {
		app.logger.Info("none of the basic rules satisfy the requirement.")

		//get final output using complex rules
		resultantMateria = useComplexRules(materia1Grade, materia2Grade, resultantMateriaGrade, materia1Type, materia2Type, *fusionReq.Materia1Mastered, *fusionReq.Materia2Mastered, &allMateria)
	} else {
		//get final output using basic rules
		resultantMateriaType := relevantBasicRule.ResultantMateriaType

		for _, materia := range allMateria {
			if materia.Grade == resultantMateriaGrade && materia.Type == string(resultantMateriaType) {
				resultantMateria.Name = materia.Name
				resultantMateria.Type = materia.DisplayType
				resultantMateria.Description = materia.Description
				break
			}
		}

	}
	app.logger.Debug("input", fusionReq.Materia1Name, fusionReq.Materia1Mastered, fusionReq.Materia2Name, fusionReq.Materia2Mastered)
	app.logger.Debug("output")
	app.logger.Debug(resultantMateria.Name)

	app.logger.Debug("end")

	err = response.JSON(w, http.StatusOK, resultantMateria)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// positions are exchanged if materia2 grade is higher
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

	if finalGrade != 8 && *req.Materia1Mastered {
		finalGrade += 1
	}
	if finalGrade != 8 && *req.Materia2Mastered {
		finalGrade += 1
	}
	return finalGrade
}

func increaseGrade(resultantMateriaGrade *int) {
	if *resultantMateriaGrade != 8 {
		*resultantMateriaGrade += 1
	}
}

func updateResultantMateriaData(allMateria *[]ccmf.Materia, resultantMateriaGrade int, resultantMateriaType string, resultantMateria *ccmf.MateriaDTO) {
	for _, materia := range *allMateria {
		if materia.Grade == resultantMateriaGrade && materia.Type == string(resultantMateriaType) {
			resultantMateria.Name = materia.Name
			resultantMateria.Type = materia.DisplayType
			resultantMateria.Description = materia.Description
			break
		}
	}
}

func useComplexRules(materia1Grade, materia2Grade, resultantMateriaGrade int, materia1Type, materia2Type string, materia1Mastered, materia2Mastered bool, allMateria *[]ccmf.Materia) (resultantMateria ccmf.MateriaDTO) {
	var resultantMateriaType string
	// Complex Rule 1: FIL, Defense
	if (materia1Type == string(ccmf.Fire) ||
		materia1Type == string(ccmf.Ice) ||
		materia1Type == string(ccmf.Lightning)) && materia2Type == string(ccmf.Defense) {
		if materia1Grade == 1 && materia2Grade == 1 {
			// output is Defense when grades are equal to 1
			resultantMateriaType = string(ccmf.Defense)
			if materia1Mastered || materia2Mastered {
				// final Grade is increased when output is Defense
				increaseGrade(&resultantMateriaGrade)
			}
		} else {
			// output is FIL when grades are NOT equal to 1
			resultantMateriaType = materia1Type
		}
		// Complex Rule 2: FIL, (Gravity, Item)
		// If materia1 is any of FIL, and materia2 is any of Gravity, Item
	} else if (materia1Type == string(ccmf.Fire) ||
		materia1Type == string(ccmf.Ice) ||
		materia1Type == string(ccmf.Lightning)) &&
		(materia2Type == string(ccmf.Gravity) ||
			materia2Type == string(ccmf.Item)) {
		if materia1Grade == materia2Grade {
			// output is Gravity / Item when grades are equal
			resultantMateriaType = materia2Type
			if materia1Mastered || materia2Mastered {
				// final Grade is increased when output is Gravity / Item
				increaseGrade(&resultantMateriaGrade)
			}
		} else {
			// output is FIL when grades are NOT equal
			resultantMateriaType = materia1Type
		}
		// Complex Rule 3: Restore, Defense
	} else if materia1Type == string(ccmf.Restore) && materia2Type == string(ccmf.Defense) {
		if (materia1Grade == 1 && materia2Grade == 1) || (materia1Grade == 4 && materia2Grade == 4) {
			resultantMateriaType = string(ccmf.Defense)
			increaseGrade(&resultantMateriaGrade)
		} else if materia2Mastered {
			resultantMateriaType = string(ccmf.Defense)
		} else {
			resultantMateriaType = string(ccmf.Restore)
		}
		// Complex Rule 4: Restore, (Gravity, Item)
		// If materia1 is Restore, and materia2 is any of Gravity, Item
	} else if materia1Type == string(ccmf.Restore) &&
		(materia2Type == string(ccmf.Gravity) ||
			materia2Type == string(ccmf.Item)) {
		if materia1Grade == 3 && materia2Grade == 3 {
			// output is Gravity / Item when grades are equal to 3
			resultantMateriaType = string(ccmf.Gravity)
			if materia1Mastered || materia2Mastered {
				// final Grade is increased when output is Gravity / Item
				increaseGrade(&resultantMateriaGrade)
			}
		} else {
			// output is Restore when grades are NOT equal
			resultantMateriaType = string(ccmf.Restore)
		}
		// Complex Rule 5: Defense, (Status Magic, FIL Status, Blade Arts Status, Quick Attack Status)
		// If materia1 is Defense, and materia2 is any of Status Magic, FIL Status, Blade Arts Status, Quick Attack Status
	} else if materia1Type == string(ccmf.Defense) &&
		(materia2Type == string(ccmf.StatusMagic) ||
			materia2Type == string(ccmf.FireStatus) ||
			materia2Type == string(ccmf.IceStatus) ||
			materia2Type == string(ccmf.LightningStatus) ||
			materia2Type == string(ccmf.BladeArtsStatus) ||
			materia2Type == string(ccmf.QuickAttackStatus)) {
		// output is always Status Defense
		resultantMateriaType = string(ccmf.StatusDefense)
		// final Grade of Status Defense is increased if input grade of materia1 is 1 or 4
		if materia1Grade == 1 || materia1Grade == 4 {
			increaseGrade(&resultantMateriaGrade)
		}
		// Complex Rule 6: Defense, Gravity
	} else if materia1Type == string(ccmf.Defense) && materia2Type == string(ccmf.Gravity) {
		if materia1Grade == 3 && materia2Grade == 3 {
			resultantMateriaType = string(ccmf.Gravity)
			increaseGrade(&resultantMateriaGrade)
		} else if materia1Grade == 7 && materia2Grade == 7 && materia2Mastered {
			resultantMateriaType = string(ccmf.Gravity)
		} else {
			resultantMateriaType = string(ccmf.Defense)
		}
		// Complex Rule 7: Defense, Item
	} else if materia1Type == string(ccmf.Defense) && materia2Type == string(ccmf.Item) {
		if materia1Grade == 7 && materia2Grade == 7 {
			resultantMateriaType = string(ccmf.Item)
		} else {
			resultantMateriaType = string(ccmf.Defense)
		}
		// Complex Rule 8: Absorb Magic, Gravity
	} else if materia1Type == string(ccmf.AbsorbMagic) && materia2Type == string(ccmf.Gravity) {
		resultantMateriaType = string(ccmf.AbsorbMagic)
		if materia1Grade == 3 && materia2Grade == 3 || materia1Grade == 5 && materia2Grade == 5 {
			increaseGrade(&resultantMateriaGrade)
		}
		// Complex Rule 9: Absorb Magic, Item
	} else if materia1Type == string(ccmf.AbsorbMagic) && materia2Type == string(ccmf.Item) {
		if materia1Grade == materia2Grade {
			resultantMateriaType = string(ccmf.Item)
			if materia1Mastered || materia2Mastered {
				increaseGrade(&resultantMateriaGrade)
			}
		} else {
			resultantMateriaType = string(ccmf.AbsorbMagic)
		}
		// Complex Rule 10: Absorb Magic, (ATKUp, VIT Up)
	} else if materia1Type == string(ccmf.AbsorbMagic) &&
		(materia2Type == string(ccmf.ATKUp) || materia2Type == string(ccmf.VITUp)) {
		resultantMateriaType = string(ccmf.AbsorbBlade)
		increaseGrade(&resultantMateriaGrade)
		// Complex Rule 11: Status Magic, Defense
	} else if materia1Type == string(ccmf.StatusMagic) && materia2Type == string(ccmf.Defense) {
		if (materia1Grade == 1 && materia2Grade == 1) || (materia1Grade == 4 && materia2Grade == 4) {
			resultantMateriaType = string(ccmf.StatusDefense)
			increaseGrade(&resultantMateriaGrade)
		} else {
			resultantMateriaType = string(ccmf.StatusMagic)
		}
		// Complex Rule 12: Status Magic, Item
	} else if materia1Type == string(ccmf.StatusMagic) && materia2Type == string(ccmf.Item) {
		if materia1Grade == materia2Grade || materia2Mastered {
			resultantMateriaType = string(ccmf.Item)
		} else {
			resultantMateriaType = string(ccmf.StatusMagic)
		}
		// Complex Rule 13: FIL Status, Defense
	} else if (materia1Type == string(ccmf.FireStatus) ||
		materia1Type == string(ccmf.IceStatus) ||
		materia1Type == string(ccmf.LightningStatus)) &&
		materia2Type == string(ccmf.Defense) {
		if materia1Grade == 1 && materia2Grade == 1 {
			resultantMateriaType = string(ccmf.StatusDefense)
			increaseGrade(&resultantMateriaGrade)
		} else {
			resultantMateriaType = materia1Type
		}
		// Complex Rule 14: FIL Status, Gravity
	} else if (materia1Type == string(ccmf.FireStatus) ||
		materia1Type == string(ccmf.IceStatus) ||
		materia1Type == string(ccmf.LightningStatus)) &&
		materia2Type == string(ccmf.Gravity) {
		if materia1Grade == 7 && materia2Grade == 7 {
			resultantMateriaType = materia1Type
		} else if materia1Grade == materia2Grade {
			resultantMateriaType = string(ccmf.Gravity)
			increaseGrade(&resultantMateriaGrade)
		} else {
			resultantMateriaType = materia1Type
		}
		// Complex Rule 15: FIL Status, Item
	} else if (materia1Type == string(ccmf.FireStatus) ||
		materia1Type == string(ccmf.IceStatus) ||
		materia1Type == string(ccmf.LightningStatus)) &&
		materia2Type == string(ccmf.Item) {
		if materia1Grade == materia2Grade {
			resultantMateriaType = string(ccmf.Item)
			increaseGrade(&resultantMateriaGrade)
		} else {
			resultantMateriaType = materia1Type
		}
		// Complex Rule 16: Gravity, (Absorb Magic, Status Magic)
	} else if materia1Type == string(ccmf.Gravity) &&
		(materia2Type == string(ccmf.AbsorbMagic) ||
			materia2Type == string(ccmf.StatusMagic)) {
		resultantMateriaType = materia2Type
		if materia1Grade == 3 || materia1Grade == 5 {
			increaseGrade(&resultantMateriaGrade)
		}
		// Complex Rule 17: Gravity, (Quick Attack, Blade Arts, FIL Blade)
	} else if materia1Type == string(ccmf.Gravity) &&
		(materia2Type == string(ccmf.QuickAttack) ||
			materia2Type == string(ccmf.BladeArts) ||
			materia2Type == string(ccmf.FireBlade) ||
			materia2Type == string(ccmf.IceBlade) ||
			materia2Type == string(ccmf.LightningBlade)) {
		resultantMateriaType = materia2Type
		if materia1Grade == 5 || materia1Mastered || materia2Mastered {
			increaseGrade(&resultantMateriaGrade)
		}

		// Complex Rule 18: Gravity, Absorb Blade
	} else if materia1Type == string(ccmf.Gravity) && materia2Type == string(ccmf.AbsorbBlade) {
		resultantMateriaType = string(ccmf.AbsorbMagic)
		if materia1Grade == 3 || materia1Grade == 5 || materia1Mastered || materia2Mastered {
			increaseGrade(&resultantMateriaGrade)
		}
		// Complex Rule 19: Gravity, Item
	} else if materia1Type == string(ccmf.Gravity) && materia2Type == string(ccmf.Item) {
		if (materia1Grade == 7 && materia2Grade == 7) || (materia1Grade == 3 && materia2Grade == 3 && materia2Mastered) {
			increaseGrade(&resultantMateriaGrade)
			resultantMateriaType = string(ccmf.Item)
		} else {
			resultantMateriaType = string(ccmf.Gravity)
		}
		if materia1Grade == 3 && materia1Mastered {
			increaseGrade(&resultantMateriaGrade)
		}
		// Complex Rule 20: Gravity, (HP Up, VIT Up, SPR Up)
	} else if materia1Type == string(ccmf.Gravity) &&
		(materia2Type == string(ccmf.HPUp) ||
			materia2Type == string(ccmf.VITUp) ||
			materia2Type == string(ccmf.SPRUp)) {
		resultantMateriaType = string(ccmf.Defense)
		if materia1Grade == 3 || materia1Mastered || materia2Mastered {
			increaseGrade(&resultantMateriaGrade)
		}

		// Complex Rule 21: Gravity, ATK Up
	} else if materia1Type == string(ccmf.Gravity) && materia2Type == string(ccmf.ATKUp) {
		resultantMateriaType = string(ccmf.BladeArts)
		if materia1Grade == 5 || materia1Mastered || materia2Mastered {
			increaseGrade(&resultantMateriaGrade)
		}
		// Complex Rule 22: Ultimate, Absorb Blade
	} else if materia1Type == string(ccmf.Ultimate) && materia2Type == string(ccmf.AbsorbBlade) {
		resultantMateriaType = string(ccmf.BladeArts)
		increaseGrade(&resultantMateriaGrade)

		// Complex Rule 23: QuickAttack, Defense
	} else if materia1Type == string(ccmf.QuickAttack) && materia2Type == string(ccmf.Defense) {
		if (materia1Grade == 1 && materia2Grade == 1) || (materia1Grade == materia2Grade && materia2Mastered) {
			resultantMateriaType = string(ccmf.Defense)
		} else {
			resultantMateriaType = string(ccmf.QuickAttack)
		}
		// Complex Rule 24: QuickAttack, Absorb Magic
	} else if materia1Type == string(ccmf.QuickAttack) && materia2Type == string(ccmf.AbsorbMagic) {
		resultantMateriaType = string(ccmf.AbsorbBlade)
		increaseGrade(&resultantMateriaGrade)

		// Complex Rule 25: QuickAttack, Gravity
	} else if materia1Type == string(ccmf.QuickAttack) && materia2Type == string(ccmf.Gravity) {
		if (materia1Grade == 3 && materia2Grade == 3) || materia1Grade == 5 && materia2Grade == 5 {
			resultantMateriaType = string(ccmf.BladeArts)
			increaseGrade(&resultantMateriaGrade)
		} else {
			resultantMateriaType = string(ccmf.QuickAttack)
		}

		// Complex Rule 26: QuickAttack, FIL Blade
	} else if materia1Type == string(ccmf.QuickAttack) &&
		(materia2Type == string(ccmf.FireBlade) ||
			materia2Type == string(ccmf.IceBlade) ||
			materia2Type == string(ccmf.LightningBlade)) {
		if materia1Grade == 7 && materia2Grade == 7 {
			resultantMateriaType = string(ccmf.QuickAttack)
		} else {
			resultantMateriaType = materia2Type
		}
		// Complex Rule 27: QuickAttack, Absorb Blade
	} else if materia1Type == string(ccmf.QuickAttack) && materia2Type == string(ccmf.AbsorbBlade) {
		resultantMateriaType = string(ccmf.AbsorbBlade)
		increaseGrade(&resultantMateriaGrade)

		// Complex Rule 28: QuickAttack, Item
	} else if materia1Type == string(ccmf.QuickAttack) && materia2Type == string(ccmf.Item) {
		if materia1Grade == materia2Grade {
			resultantMateriaType = string(ccmf.Item)
			increaseGrade(&resultantMateriaGrade)
		} else {
			resultantMateriaType = string(ccmf.QuickAttack)
		}
		// Complex Rule 29: QuickAttackStatus, Defense
	} else if materia1Type == string(ccmf.QuickAttackStatus) && materia2Type == string(ccmf.Defense) {
		if materia1Grade == 4 && materia2Grade == 4 {
			resultantMateriaType = string(ccmf.StatusDefense)
		} else {
			resultantMateriaType = string(ccmf.QuickAttackStatus)
		}

		// Complex Rule 30: QuickAttackStatus, (Absorb Magic, Absorb Blade)
	} else if materia1Type == string(ccmf.QuickAttackStatus) &&
		(materia2Type == string(ccmf.AbsorbMagic) ||
			materia2Type == string(ccmf.AbsorbBlade)) {
		resultantMateriaType = string(ccmf.AbsorbBlade)
		increaseGrade(&resultantMateriaGrade)

		// Complex Rule 31: QuickAttackStatus, Gravity
	} else if materia1Type == string(ccmf.QuickAttackStatus) && materia2Type == string(ccmf.Gravity) {
		if (materia1Grade == 5 && materia2Grade == 5) || (materia1Grade == 3 && materia2Grade == 3) {
			resultantMateriaType = string(ccmf.Gravity)
		} else {
			resultantMateriaType = string(ccmf.QuickAttackStatus)
		}
		// Complex Rule 32: QuickAttackStatus, Item
	} else if materia1Type == string(ccmf.QuickAttackStatus) && materia2Type == string(ccmf.Item) {
		if (materia1Grade == 7 && materia2Grade == 7) || (materia1Grade == 5 && materia2Grade == 5) || (materia1Grade == 3 && materia2Grade == 3) {
			resultantMateriaType = string(ccmf.Item)
			increaseGrade(&resultantMateriaGrade)
		} else {
			resultantMateriaType = string(ccmf.QuickAttackStatus)
		}
		// Complex Rule 33: Blade Arts, Defense
	} else if materia1Type == string(ccmf.BladeArts) && materia2Type == string(ccmf.Defense) {
		if materia1Grade == 1 && materia2Grade == 1 {
			resultantMateriaType = string(ccmf.Defense)
			increaseGrade(&resultantMateriaGrade)
		} else {
			resultantMateriaType = string(ccmf.BladeArts)
		}
		// Complex Rule 34: Blade Arts, Absorb Magic
	} else if materia1Type == string(ccmf.BladeArts) && materia2Type == string(ccmf.AbsorbMagic) {
		resultantMateriaType = string(ccmf.AbsorbBlade)
		increaseGrade(&resultantMateriaGrade)

		// Complex Rule 35: Blade Arts, Absorb Blade
	} else if materia1Type == string(ccmf.BladeArts) && materia2Type == string(ccmf.AbsorbBlade) {
		resultantMateriaType = string(ccmf.AbsorbBlade)
		if materia1Grade < 6 {
			increaseGrade(&resultantMateriaGrade)
		}
		// Complex Rule 36: Blade Arts, Item
	} else if materia1Type == string(ccmf.BladeArts) && materia2Type == string(ccmf.Item) {
		if (materia1Grade == 5 && materia2Grade == 5) || (materia1Grade == 3 && materia2Grade == 3) {
			resultantMateriaType = string(ccmf.Item)
			increaseGrade(&resultantMateriaGrade)
		} else {
			resultantMateriaType = string(ccmf.BladeArts)
		}
		// Complex Rule 37: Blade Arts Status, Defense
	} else if materia1Type == string(ccmf.BladeArtsStatus) && materia2Type == string(ccmf.Defense) {
		if (materia1Grade == 1 && materia2Grade == 1) || (materia1Grade == 4 && materia2Grade == 4) {
			resultantMateriaType = string(ccmf.StatusDefense)
			increaseGrade(&resultantMateriaGrade)
		} else {
			resultantMateriaType = string(ccmf.BladeArtsStatus)
		}
		// Complex Rule 38: Blade Arts Status, (Absorb Magic, Absorb Blade)
	} else if materia1Type == string(ccmf.BladeArtsStatus) &&
		(materia2Type == string(ccmf.AbsorbMagic) ||
			materia2Type == string(ccmf.AbsorbBlade)) {
		resultantMateriaType = string(ccmf.AbsorbBlade)
		increaseGrade(&resultantMateriaGrade)
		// Complex Rule 39: Blade Arts Status, Item
	} else if materia1Type == string(ccmf.BladeArtsStatus) && materia2Type == string(ccmf.Item) {
		if (materia1Grade == 7 && materia2Grade == 7) || (materia1Grade == 5 && materia2Grade == 5) {
			resultantMateriaType = string(ccmf.Item)
			increaseGrade(&resultantMateriaGrade)
		} else {
			resultantMateriaType = string(ccmf.BladeArtsStatus)
		}
		// Complex Rule 40: FIL Blade, (Restore, Defense, Status, Defense, Ultimate, QuickAttack, QuickAttackStatus, BladeArts, Punch, ATK Up, SP Turbo, Libra)
	} else if (materia1Type == string(ccmf.FireBlade) ||
		materia1Type == string(ccmf.IceBlade) ||
		materia1Type == string(ccmf.LightningBlade)) &&
		(materia2Type == string(ccmf.Restore) ||
			materia2Type == string(ccmf.StatusDefense) ||
			materia2Type == string(ccmf.Ultimate) ||
			materia2Type == string(ccmf.QuickAttack) ||
			materia2Type == string(ccmf.QuickAttackStatus) ||
			materia2Type == string(ccmf.BladeArts) ||
			materia2Type == string(ccmf.Punch) ||
			materia2Type == string(ccmf.ATKUp) ||
			materia2Type == string(ccmf.SPTurbo) ||
			materia2Type == string(ccmf.Libra)) {
		if materia1Grade == 7 {
			resultantMateriaType = string(ccmf.BladeArts)
		} else {
			resultantMateriaType = materia1Type
		}
		// Complex Rule 41: FIL Blade, Defense
	} else if (materia1Type == string(ccmf.FireBlade) ||
		materia1Type == string(ccmf.IceBlade) ||
		materia1Type == string(ccmf.LightningBlade)) && materia2Type == string(ccmf.Defense) {
		if materia1Grade == 1 {
			resultantMateriaType = string(ccmf.Defense)
		} else if materia1Grade == 7 {
			resultantMateriaType = string(ccmf.BladeArts)
		} else {
			resultantMateriaType = materia1Type
		}
		// Complex Rule 42: FIL Blade, Absorb Magic
	} else if (materia1Type == string(ccmf.FireBlade) ||
		materia1Type == string(ccmf.IceBlade) ||
		materia1Type == string(ccmf.LightningBlade)) && materia2Type == string(ccmf.AbsorbMagic) {
		if materia1Grade == 7 {
			resultantMateriaType = string(ccmf.AbsorbBlade)
		} else {
			resultantMateriaType = materia1Type
		}
		// Complex Rule 43: FIL Blade, Gravity
	} else if (materia1Type == string(ccmf.FireBlade) ||
		materia1Type == string(ccmf.IceBlade) ||
		materia1Type == string(ccmf.LightningBlade)) && materia2Type == string(ccmf.Gravity) {
		if materia1Grade == materia2Grade {
			resultantMateriaType = string(ccmf.BladeArts)
		} else {
			resultantMateriaType = string(ccmf.QuickAttack)
			increaseGrade(&resultantMateriaGrade)
		}
		// Complex Rule 44: FIL Blade, (Blade Arts Status, Absorb Blade)
	} else if (materia1Type == string(ccmf.FireBlade) ||
		materia1Type == string(ccmf.IceBlade) ||
		materia1Type == string(ccmf.LightningBlade)) &&
		(materia2Type == string(ccmf.BladeArtsStatus) ||
			materia2Type == string(ccmf.AbsorbBlade)) {
		if materia1Grade == 7 {
			resultantMateriaType = materia2Type
		} else {
			resultantMateriaType = materia1Type
			increaseGrade(&resultantMateriaGrade)
		}
		// Complex Rule 45: FIL Blade, Item
	} else if (materia1Type == string(ccmf.FireBlade) ||
		materia1Type == string(ccmf.IceBlade) ||
		materia1Type == string(ccmf.LightningBlade)) && materia2Type == string(ccmf.Item) {
		if (materia1Grade == 7 && materia2Grade == 7) || (materia1Grade == 5 && materia2Grade == 5) || (materia1Grade == 3 && materia2Grade == 3) {
			resultantMateriaType = string(ccmf.Item)
			increaseGrade(&resultantMateriaGrade)
		} else if materia1Grade == 7 && materia2Grade < 7 {
			resultantMateriaType = string(ccmf.BladeArts)
		} else {
			resultantMateriaType = string(ccmf.FireBlade)
		}
		// Complex Rule 46: FIL Blade, (HP UP, VIT UP, SPR UP)
	} else if (materia1Type == string(ccmf.FireBlade) ||
		materia1Type == string(ccmf.IceBlade) ||
		materia1Type == string(ccmf.LightningBlade)) &&
		(materia2Type == string(ccmf.HPUp) ||
			materia2Type == string(ccmf.VITUp) ||
			materia2Type == string(ccmf.SPRUp)) {
		if ((materia1Grade == 1 && materia2Grade == 1) ||
			(materia1Grade == 3 && materia2Grade == 3)) && materia2Mastered {
			resultantMateriaType = materia2Type
		} else if materia1Grade == 7 {
			resultantMateriaType = string(ccmf.BladeArts)
		} else {
			resultantMateriaType = string(ccmf.StatusDefense)
			increaseGrade(&resultantMateriaGrade)
		}
		// Complex Rule 47: Fire Blade, (MP UP, MAG Up)
	} else if materia1Type == string(ccmf.FireBlade) &&
		(materia2Type == string(ccmf.MPUp) ||
			materia2Type == string(ccmf.MAGUp)) {
		if ((materia1Grade == 1 && materia2Grade == 1) ||
			(materia1Grade == 3 && materia2Grade == 3)) && materia2Mastered {
			resultantMateriaType = materia2Type
		} else if materia1Grade == 7 {
			resultantMateriaType = string(ccmf.BladeArts)
		} else {
			resultantMateriaType = string(ccmf.FireStatus)
		}
		// Complex Rule 48: Ice Blade, (MP UP, MAG Up)
	} else if materia1Type == string(ccmf.IceBlade) &&
		(materia2Type == string(ccmf.MPUp) ||
			materia2Type == string(ccmf.MAGUp)) {
		if ((materia1Grade == 1 && materia2Grade == 1) ||
			(materia1Grade == 3 && materia2Grade == 3)) && materia2Mastered {
			resultantMateriaType = materia2Type
		} else if materia1Grade == 7 {
			resultantMateriaType = string(ccmf.BladeArts)
		} else {
			resultantMateriaType = string(ccmf.IceStatus)
		}
		// Complex Rule 49: Lightning Blade, (MP UP, MAG Up)
	} else if materia1Type == string(ccmf.LightningBlade) &&
		(materia2Type == string(ccmf.MPUp) ||
			materia2Type == string(ccmf.MAGUp)) {
		if ((materia1Grade == 1 && materia2Grade == 1) ||
			(materia1Grade == 3 && materia2Grade == 3)) && materia2Mastered {
			resultantMateriaType = materia2Type
		} else if materia1Grade == 7 {
			resultantMateriaType = string(ccmf.BladeArts)
		} else {
			resultantMateriaType = string(ccmf.LightningStatus)
		}
		// Complex Rule 50: Absorb Blade, (FIL, Restore, Defense, Status Defense, Status Magic, FIL Status, Ultimate, Quick Attack, Quick Attack Status, Blade Arts, Blade Arts Status, FIL Blade, HP, MP, AP, ATK, VIT, SP Turbo, Libra)
	} else if materia1Type == string(ccmf.AbsorbBlade) &&
		(materia2Type == string(ccmf.Fire) ||
			materia2Type == string(ccmf.Ice) ||
			materia2Type == string(ccmf.Lightning) ||
			materia2Type == string(ccmf.Restore) ||
			materia2Type == string(ccmf.Defense) ||
			materia2Type == string(ccmf.StatusDefense) ||
			materia2Type == string(ccmf.StatusMagic) ||
			materia2Type == string(ccmf.FireStatus) ||
			materia2Type == string(ccmf.IceStatus) ||
			materia2Type == string(ccmf.LightningStatus) ||
			materia2Type == string(ccmf.Ultimate) ||
			materia2Type == string(ccmf.QuickAttack) ||
			materia2Type == string(ccmf.QuickAttackStatus) ||
			materia2Type == string(ccmf.BladeArts) ||
			materia2Type == string(ccmf.BladeArtsStatus) ||
			materia2Type == string(ccmf.FireBlade) ||
			materia2Type == string(ccmf.IceBlade) ||
			materia2Type == string(ccmf.LightningBlade) ||
			materia2Type == string(ccmf.Punch) ||
			materia2Type == string(ccmf.HPUp) ||
			materia2Type == string(ccmf.MPUp) ||
			materia2Type == string(ccmf.APUp) ||
			materia2Type == string(ccmf.ATKUp) ||
			materia2Type == string(ccmf.VITUp) ||
			materia2Type == string(ccmf.SPTurbo) ||
			materia2Type == string(ccmf.Libra)) {
		resultantMateriaType = string(ccmf.AbsorbBlade)
		increaseGrade((&resultantMateriaGrade))

		// Complex Rule 51: Absorb Blade, Gravity
	} else if materia1Type == string(ccmf.AbsorbBlade) &&
		materia2Type == string(ccmf.Gravity) {
		resultantMateriaType = string(ccmf.AbsorbBlade)
		increaseGrade((&resultantMateriaGrade))
		if (materia1Grade == 3 && materia2Grade == 3) || (materia1Grade == 5 && materia2Grade == 5) {
			resultantMateriaType = string(ccmf.AbsorbMagic)
		}
		// Complex Rule 52: Absorb Blade, Item
	} else if materia1Type == string(ccmf.AbsorbBlade) &&
		materia2Type == string(ccmf.Item) {
		resultantMateriaType = string(ccmf.AbsorbBlade)
		increaseGrade((&resultantMateriaGrade))
		if (materia1Grade == 3 && materia2Grade == 3) || (materia1Grade == 5 && materia2Grade == 5) {
			resultantMateriaType = string(ccmf.Item)
		}
		// Complex Rule 53: Absorb Blade, (MAG UP, SPR UP)
	} else if materia1Type == string(ccmf.AbsorbBlade) &&
		(materia2Type == string(ccmf.MAGUp) ||
			materia2Type == string(ccmf.SPRUp)) {
		resultantMateriaType = string(ccmf.AbsorbMagic)
		if materia1Type == materia2Type && materia2Mastered {
			resultantMateriaType = materia2Type
		}

		// Complex Rule 54: Punch, FIL Blade
	} else if materia1Type == string(ccmf.Punch) &&
		(materia2Type == string(ccmf.FireBlade) ||
			materia2Type == string(ccmf.IceBlade) ||
			materia2Type == string(ccmf.LightningBlade)) {
		if materia1Grade == 7 && materia2Grade == 7 {
			resultantMateriaType = string(ccmf.BladeArts)
		} else {
			resultantMateriaType = materia2Type
		}
		// Complex Rule 55: Punch, Absorb Blade
	} else if materia1Type == string(ccmf.Punch) && materia2Type == string(ccmf.AbsorbBlade) {
		resultantMateriaType = string(ccmf.AbsorbBlade)
		increaseGrade((&resultantMateriaGrade))

		// Complex Rule 56: (HP Up, MP Up, AP Up, ATK Up, VIT Up, MAG Up, SPR Up), Defense
	} else if (materia1Type == string(ccmf.HPUp) ||
		materia1Type == string(ccmf.MPUp) ||
		materia1Type == string(ccmf.APUp) ||
		materia1Type == string(ccmf.ATKUp) ||
		materia1Type == string(ccmf.VITUp) ||
		materia1Type == string(ccmf.MAGUp) ||
		materia1Type == string(ccmf.SPRUp)) &&
		materia2Type == string(ccmf.Defense) {
		if (materia1Grade == 1 && materia2Grade == 1) || (materia1Grade == materia2Grade && materia2Mastered) {
			resultantMateriaType = string(ccmf.Defense)
			increaseGrade((&resultantMateriaGrade))
		} else {
			resultantMateriaType = materia1Type
		}
		// Complex Rule 57: (HP Up, VIT Up, SPR Up), Gravity
	} else if (materia1Type == string(ccmf.HPUp) ||
		materia1Type == string(ccmf.VITUp) ||
		materia1Type == string(ccmf.SPRUp)) &&
		materia2Type == string(ccmf.Gravity) {
		if materia1Grade == materia2Grade {
			resultantMateriaType = string(ccmf.Defense)
			increaseGrade((&resultantMateriaGrade))
		} else {
			resultantMateriaType = materia1Type
		}
		//Complex Rule 58: (MP Up, MAG Up), Gravity
	} else if (materia1Type == string(ccmf.MPUp) ||
		materia1Type == string(ccmf.MAGUp)) &&
		materia2Type == string(ccmf.Gravity) {
		if materia1Grade == materia2Grade {
			resultantMateriaType = string(ccmf.Gravity)
			increaseGrade((&resultantMateriaGrade))
		} else {
			resultantMateriaType = materia1Type
		}
		//Complex Rule 59: (AP Up), Gravity
	} else if (materia1Type == string(ccmf.APUp)) &&
		materia2Type == string(ccmf.Gravity) {
		if materia1Grade == materia2Grade {
			resultantMateriaType = string(ccmf.QuickAttack)
			increaseGrade((&resultantMateriaGrade))
		} else {
			resultantMateriaType = materia1Type
		}
		//Complex Rule 60: (ATK Up), Gravity
	} else if (materia1Type == string(ccmf.ATKUp)) &&
		materia2Type == string(ccmf.Gravity) {
		if materia1Grade == materia2Grade {
			resultantMateriaType = string(ccmf.BladeArts)
			increaseGrade((&resultantMateriaGrade))
		} else {
			resultantMateriaType = materia1Type
		}
		// Complex Rule 61: (HP Up, MP Up, AP Up, ATK Up, VIT Up, MAG Up, SPR Up), Item
	} else if (materia1Type == string(ccmf.HPUp) ||
		materia1Type == string(ccmf.MPUp) ||
		materia1Type == string(ccmf.APUp) ||
		materia1Type == string(ccmf.ATKUp) ||
		materia1Type == string(ccmf.VITUp) ||
		materia1Type == string(ccmf.MAGUp) ||
		materia1Type == string(ccmf.SPRUp)) &&
		materia2Type == string(ccmf.Item) {
		if materia1Grade == materia2Grade {
			resultantMateriaType = string(ccmf.Item)
			increaseGrade((&resultantMateriaGrade))
		} else {
			resultantMateriaType = materia1Type
		}
		// Complex Rule 62: SP Turbo, Defense
	} else if materia1Type == string(ccmf.SPTurbo) && materia2Type == string(ccmf.Defense) {
		if materia1Grade == 4 && materia2Grade == 4 {
			resultantMateriaType = string(ccmf.Defense)
		} else {
			resultantMateriaType = string(ccmf.SPTurbo)
		}
		// Complex Rule 63: SP Turbo, Gravity
	} else if materia1Type == string(ccmf.SPTurbo) && materia2Type == string(ccmf.Gravity) {
		if materia1Grade == 5 && materia2Grade == 5 {
			resultantMateriaType = string(ccmf.Gravity)
		} else {
			resultantMateriaType = string(ccmf.SPTurbo)
		}
		// Complex Rule 64: SP Turbo, Item
	} else if materia1Type == string(ccmf.SPTurbo) && materia2Type == string(ccmf.Item) {
		if (materia1Grade == 7 && materia2Grade == 7) || (materia1Grade == 5 && materia2Grade == 5) {
			resultantMateriaType = string(ccmf.Item)
		} else {
			resultantMateriaType = string(ccmf.SPTurbo)
		}
	}
	updateResultantMateriaData(allMateria, resultantMateriaGrade, resultantMateriaType, &resultantMateria)
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
			return nil, errors.New("failed to assert cached data as []Materia")
		}
	} else {
		// allMateria data is not in cache. Get from DB
		app.logger.Debug("cache miss")
		allMateria, err = app.db.GetAllMateria()
		app.setCache(string(ccmf.AllMateriaCacheKey), allMateria)
	}
	return
}
