package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/dice"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	"log"
	"math"
	"strconv"
)

func VolleyFire(col *db.Col, params map[string]interface{}) map[string]interface{} {

	adder := float64(0)

	// Go through the whole FireMods table

	firerRating := params["Rating"].(string)
	firerFatigue := params["Fatigue"].(float64)
	if firerFatigue > 4 {
		firerFatigue = 4
	}
	firerHits := params["Hits"].(float64)
	if firerHits > 6 {
		firerHits = 6
	}

	FireMods, _ := list.Get(col, "FireMod")
	for _, fmod := range FireMods.Data.([]interface{}) {
		myFireMod := fmod.(map[string]interface{})

		code := myFireMod["Code"].(string)
		val := myFireMod["Value"].(float64)
		switch code {
		case "FF":
			if params["FirstFire"].(bool) {
				adder += val
			}
		case "DIS":
			if params["Disordered"].(bool) {
				adder += val
			}
		case "SHK":
			if params["Shaken"].(bool) {
				adder += val
			}
		case "SQ":
			if params["FSquare"].(bool) {
				adder += val
			}
		case "AMD":
			if params["Ammo"].(float64) == 1 {
				adder += val
			}
		case "AME":
			if params["Ammo"].(float64) == 2 {
				adder += val
			}
		case "FLW":
			if params["LtWood"].(bool) {
				adder += val
			}
		case "FMW":
			if params["MdWood"].(bool) {
				adder += val
			}
		case "FHW":
			if params["HvWood"].(bool) {
				adder += val
			}
		case "ENL":
			if params["Enfilade"].(bool) {
				adder += val
			}
		case "TSQ":
			if params["TargetF"].(string) == "Square" {
				adder += val
			}
		case "ART":
			if params["TargetF"].(string) == "Artillery" {
				adder += val
			}
		case "TCOL":
			if params["TargetF"].(string) == "Column" {
				adder += val
			}
		case "TCC":
			if params["TargetF"].(string) == "ClosedCol" {
				adder += val
			}
		case "OO":
			if params["TargetF"].(string) == "OpenOrder" {
				adder += val
			}
		case "SK":
			if params["TargetF"].(string) == "Skirmish" {
				adder += val
			}
		case "CAV":
			if params["TargetF"].(string) == "Cavalry" {
				adder += val
			}
		case "OPP":
			if params["OppFire"].(bool) {
				adder += val
			}
		case "RN":
			if params["Rain"].(bool) {
				adder += val
			}
		case "HR":
			if params["HRain"].(bool) {
				adder += val
			}
		case "C1":
			if params["Cover"].(float64) == 1 {
				adder += val
			}
		case "C2":
			if params["Cover"].(float64) == 2 {
				adder += val
			}
		case "C3":
			if params["Cover"].(float64) == 3 {
				adder += val
			}
		case "FTG":
			adder += (val * firerFatigue)
		case "HIT":
			adder += (val * firerHits)
		default:
			if code == firerRating {
				adder += val
			}
		}
	}

	d, ammoOut := dice.ShootDice(1)
	params["EffectAmmo"] = ammoOut

	if ammoOut {
		switch params["Ammo"].(float64) {
		case 0:
			params["Ammo"] = 1
		case 1:
			params["Ammo"] = 2
		}
	}
	d2 := int(adder)
	dieScore := d + d2
	params["Dice"] = fmt.Sprintf("%d  +%d  (%d)", d, d2, dieScore)

	fid := 1

	if dieScore >= 1 {
		fid = 2
		if dieScore >= 5 {
			fid = 3
			if dieScore >= 9 {
				fid = 4
				if dieScore >= 12 {
					fid = 5
					if dieScore >= 15 {
						fid = 6
						if dieScore >= 19 {
							fid = 7
							if dieScore >= 23 {
								fid = 8
								if dieScore >= 29 {
									fid = 9
									if dieScore >= 34 {
										fid = 10
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fidString := strconv.Itoa(fid)
	FireEffectsLookup, _ := list.Lookup(col, "FireEffect", "ID")
	params["Effect"] = FireEffectsLookup[fidString]["Descr"]

	FireChartLookup, _ := list.Lookup(col, "FireChart", "ID")
	percentDamage := FireChartLookup[fidString]["Volley"].(float64)

	numBases := params["Bases"].(float64)
	damage := float64(0)
	// Hard code the values for now, assuming data for musket fire
	switch params["Range"].(float64) {
	case 0:
		damage = percentDamage * 10 * numBases
	case 1:
		damage = percentDamage * 5 * numBases
	case 2:
		damage = percentDamage * 1 * numBases
	}

	fullHits := int(math.Trunc(damage / 100))
	partialHits := int(math.Mod(damage, 100))
	extraHit := 0
	if dice.Percent(partialHits) {
		extraHit = 1
	}

	//retval["EffectHits"] = fmt.Sprintf("%f (%d, %d)", damage, fullHits, extraHit)
	params["EffectHits"] = fullHits + extraHit
	if dieScore < -5 {
		params["EffectHits"] = 0
	}

	// No longer has first fire advantage
	params["FirstFire"] = false

	return params
}

func SkirmishFire(col *db.Col, params map[string]interface{}) map[string]interface{} {

	adder := float64(0)

	// Go through the whole FireMods table
	firerRating := params["Rating"].(string)
	firerFatigue := params["Fatigue"].(float64)
	if firerFatigue > 4 {
		firerFatigue = 4
	}
	firerHits := params["Hits"].(float64)
	if firerHits > 6 {
		firerHits = 6
	}

	FireMods, _ := list.Get(col, "FireSKMod")
	for _, fmod := range FireMods.Data.([]interface{}) {
		myFireMod := fmod.(map[string]interface{})

		code := myFireMod["Code"].(string)
		val := myFireMod["Value"].(float64)
		switch code {
		case "CV":
			if params["Cover"].(bool) {
				adder += val
			}
		case "SK":
			if params["SkirmishOrder"].(bool) {
				adder += val
			}
		case "AMD":
			if params["Ammo"].(float64) == 1 {
				adder += val
			}
		case "AME":
			if params["Ammo"].(float64) == 2 {
				adder += val
			}
		case "FTG":
			adder += (val * firerFatigue)
		case "HIT":
			adder += (val * firerHits)
		default:
			if code == firerRating {
				adder += val
			}
		}
	}

	d, ammoOut := dice.ShootDice(1)
	params["EffectAmmo"] = ammoOut

	if ammoOut {
		switch params["Ammo"].(float64) {
		case 0:
			params["Ammo"] = 1
		case 1:
			params["Ammo"] = 2
		}
	}
	d2 := int(adder)
	dieScore := d + d2
	params["Dice"] = fmt.Sprintf("%d  +%d  (%d)", d, d2, dieScore)

	fid := 1

	if dieScore >= 1 {
		fid = 2
		if dieScore >= 5 {
			fid = 3
			if dieScore >= 9 {
				fid = 4
				if dieScore >= 12 {
					fid = 5
					if dieScore >= 15 {
						fid = 6
						if dieScore >= 19 {
							fid = 7
							if dieScore >= 23 {
								fid = 8
								if dieScore >= 29 {
									fid = 9
									if dieScore >= 34 {
										fid = 10
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fidString := strconv.Itoa(fid)
	FireEffectsLookup, _ := list.Lookup(col, "FireEffect", "ID")
	params["Effect"] = FireEffectsLookup[fidString]["Descr"]

	FireChartLookup, _ := list.Lookup(col, "FireChart", "ID")
	percentDamage := FireChartLookup[fidString]["SK"].(float64)

	numBases := params["Bases"].(float64)
	damage := float64(0)
	// Hard code the values for now, using less attenuation over range
	// than volley fire
	switch params["Range"].(float64) {
	case 0:
		damage = percentDamage * 6 * numBases
	case 1:
		damage = percentDamage * 4 * numBases
	case 2:
		damage = percentDamage * 2 * numBases
	}

	fullHits := int(math.Trunc(damage / 100))
	partialHits := int(math.Mod(damage, 100))
	extraHit := 0
	if dice.Percent(partialHits) {
		extraHit = 1
	}

	totalHits := fullHits + extraHit

	// Apply saving throws
	kill := 1.0
	SKEffects, _ := list.Lookup(col, "SKEffect", "ECode")
	switch params["TT"].(float64) {
	case 1:
		kill = SKEffects["C1"]["Dice"].(float64)
	case 2:
		kill = SKEffects["C2"]["Dice"].(float64)
	case 3:
		kill = SKEffects["C3"]["Dice"].(float64)
	case 4:
		kill = SKEffects["C4"]["Dice"].(float64)
	case 5:
		kill = SKEffects["C5"]["Dice"].(float64)
	}
	kill1 := int(kill)
	kill = 1.0
	switch params["TF"].(float64) {
	case 1:
		kill = SKEffects["T1"]["Dice"].(float64)
	case 2:
		kill = SKEffects["T2"]["Dice"].(float64)
	case 3:
		kill = SKEffects["T3"]["Dice"].(float64)
	case 4:
		kill = SKEffects["T4"]["Dice"].(float64)
	case 5:
		kill = SKEffects["T5"]["Dice"].(float64)
	}
	kill2 := int(kill)

	actualHits := 0
	for theHit := 1; theHit < totalHits; theHit++ {
		if dice.D6() >= kill1 {
			if dice.D6() >= kill2 {
				actualHits++
			}
		}
	}

	//retval["EffectHits"] = fmt.Sprintf("%f (%d, %d)", damage, fullHits, extraHit)
	params["EffectHits"] = totalHits
	if dieScore < -5 {
		params["EffectHits"] = 0
	}
	params["ActualHits"] = actualHits

	// No longer has first fire advantage
	params["FirstFire"] = false

	return params
}

func FireFight(col *db.Col, params map[string]interface{}) map[string]interface{} {

	adder := float64(0)

	// Go through the whole FireFightMod table
	LoserHits := params["LoserHits"].(float64)
	LHitsNow := params["LHitsNow"].(float64)
	WHitsNow := params["WHitsNow"].(float64)

	FireFightMods, _ := list.Get(col, "FireFightMod")
	for _, fmod := range FireFightMods.Data.([]interface{}) {
		myFireMod := fmod.(map[string]interface{})

		code := myFireMod["Code"].(string)
		val := myFireMod["Value"].(float64)
		switch code {
		case "AMM":
			if params["Ammo"].(float64) > 0 {
				adder += val
			}
		case "HITX":
			adder += (val * WHitsNow)
		case "HIT":
			adder += (val * LoserHits)
		case "NHIT":
			adder += (val * LHitsNow)
		case "LAV":
			switch params["LCmd"] {
			case "Poor", "Average":
				adder += val
			}
		case "LIN":
			switch params["LCmd"] {
			case "Inspirational":
				adder += val
			}
		case "LCH":
			switch params["LCmd"] {
			case "Charismatic":
				adder += val
			}
		}
	}

	d := dice.DieRoll()
	d2 := int(adder)
	dieScore := d + d2
	params["Dice"] = fmt.Sprintf("%d  +%d  (%d)", d, d2, dieScore)

	fid := 2

	if dieScore >= 13 {
		fid = 13
		if dieScore >= 17 {
			fid = 17
			if dieScore >= 20 {
				fid = 20
				if dieScore >= 23 {
					fid = 23
				}
			}
		}
	}

	fidString := strconv.Itoa(fid)
	FireFightLookup, _ := list.Lookup(col, "FireFight", "Dice")
	FireFight := FireFightLookup[fidString]

	params["Result"] = FireFight["Descr"]
	params["FallBack"] = FireFight["FallBack"]
	params["HoldCover"] = FireFight["HoldCover"]
	params["Disorder"] = FireFight["Disorder"]
	params["Rout"] = FireFight["Rout"]

	if FireFight["HoldCover"].(bool) && params["Cover"].(bool) {
		params["FallBack"] = false
	}
	return params
}

func ArtyFire(col *db.Col, params map[string]interface{}) map[string]interface{} {

	adder := float64(0)

	// Go through the whole ArtMod table

	/*
		ArtMods, _ := list.Get(col, "ArtMod")
		for _, amod := range ArtMods.Data.([]interface{}) {
			myArtMod := amod.(map[string]interface{})

			code := myArtMod["Code"].(string)
			val := myArtMod["Value"].(float64)
			switch code {
			}
		}
	*/
	d := dice.DieRoll()
	d2 := int(adder)
	dieScore := d + d2
	params["Dice"] = fmt.Sprintf("%d  +%d  (%d)", d, d2, dieScore)

	fid := 1

	if dieScore >= 1 {
		fid = 2
		if dieScore >= 5 {
			fid = 3
			if dieScore >= 9 {
				fid = 4
				if dieScore >= 12 {
					fid = 5
					if dieScore >= 15 {
						fid = 6
						if dieScore >= 19 {
							fid = 7
							if dieScore >= 23 {
								fid = 8
								if dieScore >= 29 {
									fid = 9
									if dieScore >= 34 {
										fid = 10
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fidString := strconv.Itoa(fid)
	log.Println(fidString)

	return params
}
