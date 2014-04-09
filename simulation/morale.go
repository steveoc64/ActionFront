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
