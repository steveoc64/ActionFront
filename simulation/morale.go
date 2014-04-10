package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/dice"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	"log"
)

func UnitMoraleTest(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Cover := params["Cover"].(float64)
	Enfilade := params["Enfilade"].(float64)
	Hits := params["Hits"].(float64)
	Fatigue := params["Fatigue"].(float64)
	Bases := params["Bases"].(float64)
	Formation := params["Formation"].(string)
	Rating := params["Rating"].(string)
	Leader := params["Leader"].(string)
	params["Dice"] = ""
	params["Effect"] = ""
	params["PassScore"] = ""

	adder := float64(0)

	UnitMoraleTest := list.Lookup(col, "UnitMoraleTest", "Rating")
	PassScore := int(UnitMoraleTest[Rating]["Pass"].(float64))

	// Go through the whole UnitMoraleMod table

	UnitMoraleMods, _ := list.Get(col, "UnitMoraleMod")
	for _, mod := range UnitMoraleMods.Data.([]interface{}) {
		myUnitMoraleMod := mod.(map[string]interface{})

		code := myUnitMoraleMod["Code"].(string)
		val := myUnitMoraleMod["Value"].(float64)
		switch code {
		case "C1": // Light Cover
			if Cover == 1 {
				adder += val
				log.Println(code, val)
			}
		case "C2": // Medium Cover
			if Cover == 2 {
				adder += val
				log.Println(code, val)
			}
		case "C3": // Heavy Cover
			if Cover == 3 {
				adder += val
				log.Println(code, val)
			}
		case "C4": // Super Cover
			if Cover == 4 {
				adder += val
				log.Println(code, val)
			}
		case "F1": // Enfiladed by Infantry point blank
			if Enfilade == 1 {
				adder += val
				log.Println(code, val)
			}
		case "F2": // Enfiladed by infantry close
			if Enfilade == 2 {
				adder += val
				log.Println(code, val)
			}
		case "F3": // Enfiladed by artillery
			if Enfilade == 3 {
				adder += val
				log.Println(code, val)
			}
		case "DIS": // Disordered
			if params["Disordered"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "GC": // Charged by Guard
			if params["GCharge"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "KL": // Unformed - form klumpen
			if Formation == "OO" {
				adder += val
				log.Println(code, val)
			}
		case "HW": // Heavy Woods
			if params["HvWoods"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "CX": // Caisson explodes
			if params["CX"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "BB": // Bombardment only
			if params["BBOnly"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "SQ": // Square
			if Formation == "SQ" {
				adder += val
				log.Println(code, val)
			}
		case "CC": // Closed Col
			if Formation == "CC" {
				adder += val
				log.Println(code, val)
			}
		case "L1": // Veteran in Line
			if Formation == "Line" && Rating == "Veteran" {
				adder += val
				log.Println(code, val)
			}
		case "L2": // Regular in Line
			if Formation == "Line" && Rating == "Regular" {
				adder += val
				log.Println(code, val)
			}
		case "L3": // Conscript or lower in Line
			if Formation == "Line" {
				switch Rating {
				case "Conscript", "Landwehr", "Militia", "Rabble":
					if Bases > 2 {
						adder += val * Bases
					} else {
						adder += -2
					}
					log.Println(code, val)
				}
			}

		case "HIT": // Hits
			adder += val * Hits
		case "FT": // Hits
			adder += val * Fatigue
		}
	}

	switch Leader {
	case "Despicable":
		adder += -2
	case "Poor":
		adder += -1
	case "Inspirational":
		adder += 1
	case "Charismatic":
		adder += 3
	}

	d := dice.DieRoll()
	dieScore := d + int(adder)
	params["Dice"] = fmt.Sprintf("%d + %d (%d)", d, int(adder), dieScore)
	params["PassScore"] = PassScore - int(adder)

	if dieScore > PassScore {
		params["Effect"] = "Steady Morale"
	} else if dieScore > PassScore-1 {
		params["Effect"] = "Shaken"

	} else {
		params["Effect"] = "Morale Broken"
	}

	return params
}

func MEMoraleTest(col *db.Col, params map[string]interface{}) map[string]interface{} {

	params["Dice"] = ""
	params["Effect"] = ""
	Fatigue := params["Fatigue"].(float64)
	CFatigue := params["CFatigue"].(float64)
	GT := params["GT"].(float64)
	Leader := params["Leader"].(string)
	params["PassScore"] = ""
	params["Dice"] = ""
	params["Effect"] = ""
	params["EffectSteady"] = false
	params["EffectShaken"] = false
	params["EffectRetreat"] = false
	params["EffectBroken"] = false
	params["EffectFatigue"] = 0

	adder := float64(0)

	// Go through the whole MEMoraleMod table
	MEMoraleMods, _ := list.Get(col, "MEMoraleMod")
	for _, mod := range MEMoraleMods.Data.([]interface{}) {
		myMEMoraleMod := mod.(map[string]interface{})

		code := myMEMoraleMod["Code"].(string)
		val := myMEMoraleMod["Value"].(float64)
		switch code {
		case "BADI": // Infantry in BAD morale
			adder += val * params["BadI"].(float64)
			log.Println(code, val)
		case "BADA": // Artillery in BAD morale
			adder += val * params["BadA"].(float64)
			log.Println(code, val)
		case "BADC": // Cavalry in BAD morale
			adder += val * params["BadC"].(float64)
			log.Println(code, val)
		case "GOOD": // Any unit in good morale, with full ammo
			adder += val * params["Good"].(float64)
			log.Println(code, val)
		case "CAW": // Close action victories this turn
			adder += val * params["CAW"].(float64)
			log.Println(code, val)
		case "CAD": // Close action losses this turn
			adder += val * params["CAD"].(float64)
			log.Println(code, val)
		case "Fatigue": // Fatigue level
			adder += val * Fatigue
			log.Println(code, val)
		case "SPH": // Per structure still held
			adder += val * params["Sown"].(float64)
			log.Println(code, val)
		case "SPL": // Per structure lost to the enemy
			adder += val * params["SE"].(float64)
			log.Println(code, val)
		case "CF1": // Campaign fatigue weary
			if CFatigue == 1 {
				adder += val
				log.Println(code, val)
			}
		case "CF2": // Campaign fatigue haggard
			if CFatigue == 2 {
				adder += val
				log.Println(code, val)
			}
		case "CF3": // Campaign fatigue spent
			if CFatigue == 3 {
				adder += val
				log.Println(code, val)
			}
		case "INTER": // Interpenetrated
			if params["Interp"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "SHK": // Already shaken
			if params["PrevSH"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "SQP": // Adjacent ME SPQ'd
			if params["SPQ"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "SP": // Enemy strongpoint has fallen
			if params["ESP"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "COLD": // Cold weather
			if params["Cold"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "F1": // GT Flanked - deployed
			if GT == 1 {
				adder += val
				log.Println(code, val)
			}
		case "F2": // GT Flanked - cond col
			if GT == 2 {
				adder += val
				log.Println(code, val)
			}
		case "F3": // GT Flanked - reg col
			if GT == 3 {
				adder += val
				log.Println(code, val)
			}
		case "F4": // GT Flanked - extd col
			if GT == 4 {
				adder += val
				log.Println(code, val)
			}
		}
	}

	switch Leader {
	case "Despicable":
		adder += -2
	case "Poor":
		adder += -1
	case "Inspirational":
		adder += 1
	case "Charismatic":
		adder += 3
	}

	d := dice.DieRoll()
	dieScore := d + int(adder)
	params["Dice"] = fmt.Sprintf("%d + %d (%d)", d, int(adder), dieScore)

	addFatigue := float64(0)
	if dieScore >= 0 {
		params["Effect"] = "ME Breaks, retreat 2Grids in BAD morale"
		params["EffectSteady"] = false
		params["EffectShaken"] = true
		params["EffectRetreat"] = true
		params["EffectBroken"] = true
		addFatigue = 1

		if dieScore >= 6 {
			params["Effect"] = "ME Retreats 2Grids Shaken, convert to Break Off order"
			params["EffectSteady"] = false
			params["EffectShaken"] = true
			params["EffectRetreat"] = true
			params["EffectBroken"] = false
			addFatigue = 1

			if dieScore >= 9 {
				params["Effect"] = "ME becomes Shaken. Attacks without Impetus fall back 2Grids, revert to Defend order"
				params["EffectSteady"] = false
				params["EffectShaken"] = true
				params["EffectRetreat"] = false
				params["EffectBroken"] = false
				addFatigue = 0

				if dieScore >= 11 {
					params["Effect"] = "ME Remains Steady"
					params["EffectSteady"] = true
					params["EffectShaken"] = false
					params["EffectRetreat"] = false
					params["EffectBroken"] = false
					addFatigue = 0
				}
			}
		}
	}
	newFatigue := Fatigue + addFatigue
	if newFatigue > 4 {
		newFatigue = 4
	}
	params["Fatigue"] = newFatigue
	params["EffectFatigue"] = addFatigue
	if params["EffectShaken"].(bool) {
		params["PrevSH"] = true
	}
	if params["EffectSteady"].(bool) {
		params["PrevSH"] = false
	}

	return params
}

func MEPanicTest(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Cover := params["Cover"].(float64)
	Enfilade := params["Enfilade"].(float64)
	Hits := params["Hits"].(float64)
	Fatigue := params["Fatigue"].(float64)
	Bases := params["Bases"].(float64)
	Formation := params["Formation"].(string)
	Rating := params["Rating"].(string)
	Leader := params["Leader"].(string)
	params["Dice"] = ""
	params["Effect"] = ""

	adder := float64(0)

	UnitMoraleTest := list.Lookup(col, "UnitMoraleTest", "Rating")
	PassScore := int(UnitMoraleTest[Rating]["Pass"].(float64))

	// Go through the whole UnitMoraleMod table

	UnitMoraleMods, _ := list.Get(col, "UnitMoraleMod")
	for _, mod := range UnitMoraleMods.Data.([]interface{}) {
		myUnitMoraleMod := mod.(map[string]interface{})

		code := myUnitMoraleMod["Code"].(string)
		val := myUnitMoraleMod["Value"].(float64)
		switch code {
		case "C1": // Light Cover
			if Cover == 1 {
				adder += val
				log.Println(code, val)
			}
		case "C2": // Medium Cover
			if Cover == 2 {
				adder += val
				log.Println(code, val)
			}
		case "C3": // Heavy Cover
			if Cover == 3 {
				adder += val
				log.Println(code, val)
			}
		case "C4": // Super Cover
			if Cover == 4 {
				adder += val
				log.Println(code, val)
			}
		case "F1": // Enfiladed by Infantry point blank
			if Enfilade == 1 {
				adder += val
				log.Println(code, val)
			}
		case "F2": // Enfiladed by infantry close
			if Enfilade == 2 {
				adder += val
				log.Println(code, val)
			}
		case "F3": // Enfiladed by artillery
			if Enfilade == 3 {
				adder += val
				log.Println(code, val)
			}
		case "DIS": // Disordered
			if params["Disordered"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "GC": // Charged by Guard
			if params["GCharge"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "KL": // Unformed - form klumpen
			if Formation == "OO" {
				adder += val
				log.Println(code, val)
			}
		case "HW": // Heavy Woods
			if params["HvWoods"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "CX": // Caisson explodes
			if params["CX"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "BB": // Bombardment only
			if params["BBOnly"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "SQ": // Square
			if Formation == "SQ" {
				adder += val
				log.Println(code, val)
			}
		case "CC": // Closed Col
			if Formation == "CC" {
				adder += val
				log.Println(code, val)
			}
		case "L1": // Veteran in Line
			if Formation == "Line" && Rating == "Veteran" {
				adder += val
				log.Println(code, val)
			}
		case "L2": // Regular in Line
			if Formation == "Line" && Rating == "Regular" {
				adder += val
				log.Println(code, val)
			}
		case "L3": // Conscript or lower in Line
			if Formation == "Line" {
				switch Rating {
				case "Conscript", "Landwehr", "Militia", "Rabble":
					if Bases > 2 {
						adder += val * Bases
					} else {
						adder += -2
					}
					log.Println(code, val)
				}
			}

		case "HIT": // Hits
			adder += val * Hits
		case "FT": // Hits
			adder += val * Fatigue
		}
	}

	switch Leader {
	case "Despicable":
		adder += -2
	case "Poor":
		adder += -1
	case "Inspirational":
		adder += 1
	case "Charismatic":
		adder += 3
	}

	d := dice.DieRoll()
	dieScore := d + int(adder)
	params["Dice"] = fmt.Sprintf("%d + %d (%d)", d, int(adder), dieScore)
	params["PassScore"] = PassScore - int(adder)

	if dieScore > PassScore {
		params["Effect"] = "Steady Morale"
	} else if dieScore > PassScore-1 {
		params["Effect"] = "Shaken"

	} else {
		params["Effect"] = "Morale Broken"
	}

	return params
}
