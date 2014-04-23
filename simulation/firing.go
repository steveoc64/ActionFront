package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/dice"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	"log"
	"math"
	"strconv"
	"strings"
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
		case "C4":
			if params["Cover"].(float64) == 4 {
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
	if params["FirstFire"].(bool) {
		ammoOut = false
	}
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
	FireEffectsLookup := list.Lookup(col, "FireEffect", "ID")
	params["Effect"] = FireEffectsLookup[fidString]["Descr"]

	FireChartLookup := list.Lookup(col, "FireChart", "ID")
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
	FireEffectsLookup := list.Lookup(col, "FireEffect", "ID")
	params["Effect"] = FireEffectsLookup[fidString]["Descr"]

	FireChartLookup := list.Lookup(col, "FireChart", "ID")
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
	SKEffects := list.Lookup(col, "SKEffect", "ECode")
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
	FireFightLookup := list.Lookup(col, "FireFight", "Dice")
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

// Calculate the fire from artillery
func ArtyFire(col *db.Col, params map[string]interface{}) map[string]interface{} {

	adder := float64(0)

	TFormation := params["TFormation"].(string)
	Angle := params["Angle"].(float64)
	TType := params["TType"].(string)
	isInf := TType == "Infantry"
	isCav := TType == "Cavalry"
	GunneryClass := params["GunneryClass"].(float64)
	Range := params["Range"].(float64)
	ArtyWeight := params["ArtyWeight"].(string)

	ArtyRanges := list.Lookup(col, "ArtRange", "Weight")
	GunRange := ArtyRanges[ArtyWeight]
	ShortRange := GunRange["Short"].(float64)
	MediumRange := GunRange["Medium"].(float64)
	LongRange := GunRange["Long"].(float64)
	MaxRange := GunRange["Max"].(float64)

	Ammo := params["Ammo"].(float64)
	FireMission := params["FireMission"].(string)
	NumBases := params["Bases"].(float64)
	Fatigue := params["Fatigue"].(float64)
	Followup := params["Followup"].(bool)

	params["Hits"] = 0
	params["Effect"] = ""
	params["EffectRetire"] = false

	// Apply some common sense rules to ranges and fire missions
	ammoOutScore := 2

	switch FireMission {
	case "Bombardment":
		if Range > MaxRange {
			params["Effect"] = fmt.Sprintf("Out of Range - Max Range = %dm (%d Grids)", int(MaxRange*400), int(MaxRange))
			return params
		}
		if Fatigue >= 4 {
			params["Effect"] = fmt.Sprintf("Bty is exhausted")
			params["EffectRetire"] = true
			return params
		}
	case "Tactical":
		if Range > LongRange {
			params["Effect"] = fmt.Sprintf("Too Far for tactical fire !! Max Effective Range = %dm (%d Grids)", int(LongRange*400), int(LongRange))
			return params
		}
		if Fatigue >= 4 {
			params["Effect"] = fmt.Sprintf("Bty is exhausted")
			params["EffectRetire"] = true
			return params
		}
	case "DefBty":
		if Range > ShortRange {
			params["Effect"] = fmt.Sprintf("Defence of Bty only applies to Short Range fire at %dm (%d Grids)", int(ShortRange*400), int(ShortRange))
			return params
		}
	case "Support":
		if Range > MediumRange {
			params["Effect"] = fmt.Sprintf("Support Fire max range = %dm (%d Grids)", int(MediumRange*400), int(MediumRange))
			return params
		}
		if Fatigue >= 4 {
			params["Effect"] = fmt.Sprintf("Bty is exhausted")
			params["EffectRetire"] = true
			return params
		}
	case "FireRetire":
		if Range > MediumRange {
			params["Effect"] = fmt.Sprintf("Fire and Retire max range = %dm (%d Grids)", int(MediumRange*400), int(MediumRange))
			return params
		}
		if Fatigue >= 4 {
			params["Effect"] = fmt.Sprintf("Bty is exhausted - will stand in defence of Bty instead")
			params["EffectRetire"] = false
			return params
		}
	}
	params["Effect"] = ""

	// Go through the whole ArtMod table

	ArtMods, _ := list.Get(col, "ArtMod")
	for _, amod := range ArtMods.Data.([]interface{}) {
		myArtMod := amod.(map[string]interface{})

		code := myArtMod["Code"].(string)
		val := myArtMod["Value"].(float64)
		switch code {
		case "MO": // Mixed order target
			if isInf && TFormation == "MO" {
				adder += val
				log.Println(code, val)
			}
		case "SQ": // Target in Square
			if isInf && TFormation == "SQ" {
				adder += val
				log.Println(code, val)
			}
		case "RS": // Target on reverse slope
			if params["ReverseSlope"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "CAV": // Target is cavalry
			if TType == "CV" {
				adder += val
				log.Println(code, val)
			}
		case "CW": // Withdrawing Cavalry
			if TType == "CW" {
				adder += val
				log.Println(code, val)
			}
		case "CC": // Closed Column
			if isInf && TFormation == "CC" {
				adder += val
				log.Println(code, val)
			}
		case "CO": // Column formation
			if isInf && TFormation == "CO" {
				adder += val
				log.Println(code, val)
			}
		case "OO": // Open Order
			if isInf && TFormation == "OO" {
				adder += val
				log.Println(code, val)
			}
		case "G0": // Grade Old Guard
			if GunneryClass == 0 {
				adder += val
				log.Println(code, val)
			}
		case "G1": // Class I
			if GunneryClass == 1 {
				adder += val
				log.Println(code, val)
			}
		case "G2": // Class II
			if GunneryClass == 2 {
				adder += val
				log.Println(code, val)
			}
		case "G3": // Class III
			if GunneryClass == 3 {
				adder += val
				log.Println(code, val)
			}
		case "RM": // Medium Range - which is 2 grids for everything
			if Range == MediumRange {
				adder += val
				log.Println(code, val)
			}
		case "RL": // Long Range - which is dependant on the type of gun
			if Range > MediumRange && Range <= LongRange {
				adder += val
				log.Println(code, val)
			}
		case "SC": // Screened by Skirmishers
			if params["Screened"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "NAP": // Napoleon Himself attached to battery
			if params["NapAttached"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "CR": // Corps Commander Attached to Bty
			if params["CCAttached"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "RE": // Fresh out of Reserve
			if Ammo == 0 {
				adder += val
				log.Println(code, val)
			}
		case "FT": // Per Fatigue Level
			adder += val * Fatigue
			log.Println(code, val)
		case "HC": // Heavy Guns, counter bty fire at long range
			if TType == "CounterBty" && Range > MediumRange {
				switch ArtyWeight {
				case "Heavy", "MdHeavy":
					adder += val
					log.Println(code, val)
				}
			}
		case "LC": // Light or medium guns, counter bty fire at long range
			if TType == "CounterBty" && Range > MediumRange {
				switch ArtyWeight {
				case "Light", "Medium":
					adder += val
					log.Println(code, val)
				}
			}
		case "A1": // Ammo Depleted
			if Ammo == 2 {
				adder += val
				log.Println(code, val)
			}
		case "A2": // Ammo exhd
			if Ammo == 3 {
				adder += val
				log.Println(code, val)
			}
		case "T1": // Marchfeld terrain
			if params["Marchfeld"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "HR": // Heavy Rain
			if params["HvRain"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "3G": // 3Gun sections
			if params["ThreeGun"].(bool) {
				adder += val
				log.Println(code, val)
			}
		case "S1": // Def bty vs Infantry
			if isInf && !Followup && FireMission == "DefBty" {
				adder += val
				log.Println(code, val)
			}
		case "S2": // Fire and retire vs infantry
			if isInf && !Followup && FireMission == "FireRetire" {
				adder += val
				log.Println(code, val)
			}
		case "S3": // Def Bty followup vs Infantry
			if isInf && Followup && FireMission == "DefBty" {
				adder += val
				log.Println(code, val)
			}
		case "S4": // Followup fire and retire vs Infantry
			if isInf && Followup && FireMission == "FireRetire" {
				adder += val
				log.Println(code, val)
			}
		case "S5": // Support fire vs Infantry
			if isInf && !Followup && FireMission == "Support" {
				adder += val
				log.Println(code, val)
			}
		case "S6": // Followup support vs Infantry
			if isInf && Followup && FireMission == "Support" {
				adder += val
				log.Println(code, val)
			}
		case "S7": // Defence of Bty vs Cav
			if isCav && !Followup && FireMission == "DefBty" {
				adder += val
				log.Println(code, val)
			}
		case "S8": // Fire & Retire vs Cav
			if isCav && !Followup && FireMission == "FireRetire" {
				adder += val
				log.Println(code, val)
			}
		case "S9": // Followup Defence of Bty vs Cav
			if isCav && Followup && FireMission == "DefBty" {
				adder += val
				log.Println(code, val)
			}
		case "S10": // Followup fire & retire vs Cav
			if isCav && Followup && FireMission == "FireRetire" {
				adder += val
				log.Println(code, val)
			}
		case "S11": // Support Fire vs Cav
			if isCav && !Followup && FireMission == "Support" {
				adder += val
				log.Println(code, val)
			}
		case "S12": // Followup Support Fire vs Cav
			if isCav && Followup && FireMission == "Support" {
				adder += val
				log.Println(code, val)
			}
		case "B1":
		case "B2":
		case "B3":
		case "B4":
		case "B5":

		}
	}
	/*
		gameData.Insert(DataMap("ArtMod", ArtMod{"B1", "Bombardment - Grand Bty with rated leader", 5}))
		gameData.Insert(DataMap("ArtMod", ArtMod{"B2", "Bombardment - Grand Bty without rated leader", 3}))
		gameData.Insert(DataMap("ArtMod", ArtMod{"B3", "Bombardment - Short Range Fire", 3}))
		gameData.Insert(DataMap("ArtMod", ArtMod{"B4", "Bombardment - Paced Rate of Fire", -4}))
		gameData.Insert(DataMap("ArtMod", ArtMod{"B5", "Bombardment - Intensive Rate of Fire", 6}))
		gameData.Insert(DataMap("ArtMod", ArtMod{"B5", "Bombardment - Target Moved 1 Grid", -7}))
		gameData.Insert(DataMap("ArtMod", ArtMod{"B5", "Bombardment - Target Moved more than 1 Grid", -10}))
	*/
	params["EffectAmmo"] = false
	d, ammoOut := dice.ShootDice(ammoOutScore)
	if ammoOut {
		params["EffectAmmo"] = true
		switch Ammo {
		case 0:
			params["EffectAmmo"] = false
		case 1:
			params["Ammo"] = 2
		case 2:
			params["Ammo"] = 3
		}
	}
	// If First Fire out of reserve, then advance to full ammo status
	if Ammo == 0 {
		params["Ammo"] = 1
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
	FireEffectsLookup := list.Lookup(col, "FireEffect", "ID")
	params["Effect"] = FireEffectsLookup[fidString]["Descr"]

	FireChartLookup := list.Lookup(col, "FireChart", "ID")
	scoreToHit := float64(0)
	switch ArtyWeight {
	case "Light":
		scoreToHit = FireChartLookup[fidString]["LtArt"].(float64)
	case "Medium":
		scoreToHit = FireChartLookup[fidString]["MdArt"].(float64)
	case "MdHeavy":
		scoreToHit = FireChartLookup[fidString]["MdHvArt"].(float64)
	case "Heavy":
		scoreToHit = FireChartLookup[fidString]["HvArt"].(float64)
	}
	log.Println("Score to hit on D12 = ", scoreToHit)

	AngleMultiplier := 1.0
	switch Angle {
	case 1:
		AngleMultiplier = 1.5
	case 2:
		AngleMultiplier = 2.0
	}

	numHits := dice.BucketD12(int(NumBases*4), int(scoreToHit))
	log.Println(numHits, "hits, mult=", AngleMultiplier)
	params["Hits"] = math.Trunc(AngleMultiplier * float64(numHits))

	Fatigue++
	if Fatigue > 4 {
		Fatigue = 4
	}
	params["Fatigue"] = Fatigue

	return params
}

// Calculate the hits on artillery
func CounterBty(col *db.Col, params map[string]interface{}) map[string]interface{} {

	adder := float64(0)
	Deploy := int(params["Deploy"].(float64))
	Hits := int(params["Hits"].(float64))
	Cover := params["Cover"].(bool)
	Shrapnel := params["Shrapnel"].(bool)

	if Shrapnel {
		adder += 1
	}

	DiceResults := make([]string, Hits)
	CrewHits := 0
	HorseHits := 0
	CaissonHits := false

	CounterBty := list.Lookup(col, "CounterBty", "Score")

	for i := 0; i < Hits; i++ {
		// Roll the Dice
		Dice := dice.DieRoll()
		TotalDice := Dice + int(adder)
		DiceResults[i] = fmt.Sprintf("%d", Dice)

		fid := ""
		if TotalDice >= 2 {
			fid = "2"
			if TotalDice >= 9 {
				fid = "9"
				if TotalDice >= 12 {
					fid = "12"
					if TotalDice >= 14 {
						fid = "14"
						if TotalDice >= 19 {
							fid = "19"
						}
					}
				}
			}
		}

		CBRecord := CounterBty[fid]
		Horsery := 0
		if Deploy == 0 {
			CrewHits += int(CBRecord["Crew"].(float64))
			Horsery = int(CBRecord["Horses"].(float64))
			if Cover && Horsery > 1 {
				Horsery--
			}
			HorseHits += Horsery
			if CBRecord["Caisson"].(bool) {
				CaissonHits = true
			}
		} else {
			CrewHits += int(CBRecord["LCrew"].(float64))
			Horsery = int(CBRecord["LHorses"].(float64))
			if Cover && Horsery > 1 {
				Horsery--
			}
			HorseHits += Horsery
			if CBRecord["Caisson"].(bool) {
				CaissonHits = true
			}
		}
	}
	params["Dice"] = strings.Join(DiceResults, ", ")
	params["ResultCrew"] = fmt.Sprintf("%d Crew Hits", CrewHits)
	params["ResultHorse"] = fmt.Sprintf("%d Horse Hits", HorseHits)
	if CaissonHits {
		params["ResultCaisson"] = "Explodes !"
	} else {
		params["ResultCaisson"] = ""
	}

	return params
}
