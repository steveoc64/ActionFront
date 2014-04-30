package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/dice"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	"log"
	"math"
)

// Check Unit Morale
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

// If an ME breaks, test adjacent MEs from right to left, provided that they are within 2 grids of the failing ME
func MEPanicTest(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Rating := params["Rating"].(string)
	Fatigue := params["Fatigue"].(float64)
	CFatigue := params["CFatigue"].(float64)
	Status := params["Status"].(float64)
	Sown := params["Sown"].(float64)
	OGS := params["OGS"].(bool)
	OGB := params["OGB"].(bool)
	Interp := params["Interp"].(bool)
	PrevSH := params["PrevSH"].(bool)
	TRAP := params["TRAP"].(bool)
	WITH := params["WITH"].(bool)

	params["Dice"] = ""
	params["PassScore"] = ""
	params["Effect"] = ""

	adder := float64(0)

	MEPanicTest := list.Lookup(col, "MEPanicTest", "Rating")
	PassScore := int(MEPanicTest[Rating]["CarryOn"].(float64))
	BrokenScore := int(MEPanicTest[Rating]["Broken"].(float64))
	ShakenScore := int(MEPanicTest[Rating]["Shaken"].(float64))
	params["PassScore"] = PassScore

	// Go through the whole UnitMoraleMod table

	Mods, _ := list.Get(col, "MEPanicMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "25":
			if Status == 1 {
				adder += val
			}
		case "50":
			if Status == 2 {
				adder += val
			}
		case "CF1":
			if CFatigue == 1 {
				adder += val
			}
		case "CF2":
			if CFatigue == 2 {
				adder += val
			}
		case "CF3":
			if CFatigue == 3 {
				adder += val
			}
		case "Fatigue":
			adder += val * Fatigue
		case "GOOD":
			if Status == 0 {
				adder += val
			}
		case "INTER":
			if Interp {
				adder += val
			}
		case "OG1":
			if OGS {
				adder += val
			}
		case "OG2":
			if OGB {
				adder += val
			}
		case "SHK":
			if PrevSH {
				adder += val
			}
		case "SP":
			adder += val * Sown
		case "TRAP":
			if TRAP {
				adder += val
			}
		case "WTH":
			if WITH {
				adder += val
			}
		}
	}

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	params["Broken"] = false
	params["ResultShaken"] = false
	if TotalDice <= BrokenScore {
		params["Effect"] = "Panic - Entire ME dissolves in Bad Morale, and retires 2 grids"
		Fatigue++
		if Fatigue > 4 {
			Fatigue = 4
		}
		params["Fatigue"] = Fatigue
		params["ResultBroken"] = true
	} else if TotalDice <= ShakenScore {
		params["Effect"] = "Shaken - ME is Shaken"
		params["ResultShaken"] = true
		params["PrevSH"] = true
	} else {
		params["Effect"] = "ME will Carry On in good order"
	}

	return params
}

// Recovery from Bad Morale
func BadMoraleRec(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Rating := params["Rating"].(string)
	Leader := params["Leader"].(string)
	METype := params["METype"].(float64)
	Hits := params["Hits"].(float64)
	Fatigue := params["Fatigue"].(float64)
	LostStandard := params["LostStandard"].(bool)

	adder := float64(0)
	Mods, _ := list.Get(col, "BUAMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "SL":
			if LostStandard {
				adder += val
			}
		case "HIT":
			adder += Hits * val
		case "CF":
			if METype == 2 {
				adder += Fatigue * val
			}
		case "AF":
			if METype == 3 {
				adder += Fatigue * val
			}
		case "MF":
			if METype == 0 || METype == 1 {
				adder += Fatigue * val
			}
		}
	}

	switch Leader {
	case "UnInspiring":
		adder += -1
	case "Average":
		adder += 0
	case "Inspirational":
		adder += 1
	case "Charismatic":
		adder += 3
	}

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	// Get the BadMorale recovery table
	BadMoraleRec := list.Lookup(col, "BadMoraleRec", "Rating")[Rating]
	GoodMoraleScore := int(BadMoraleRec["GoodMorale"].(float64))
	TryAgainScore := int(BadMoraleRec["TryAgain"].(float64))
	if TotalDice >= GoodMoraleScore {
		params["Result"] = "Unit Rallies, and is ready for Battle"
		params["ResultSteady"] = true
		params["ResultContinue"] = false
		params["ResultLeaves"] = false
		// Get some hits back
		Hits = math.Trunc(Hits / 2)
		if Fatigue > 0 {
			Fatigue--
		}
		params["Hits"] = Hits
		params["Fatigue"] = Fatigue
	} else if TotalDice >= TryAgainScore {
		params["Result"] = "Unit continues to rally - carry on and try again next turn"
		params["ResultSteady"] = false
		params["ResultContinue"] = true
		params["ResultLeaves"] = false
	} else {
		params["Result"] = "Unit has lost confidence, and leaves the field for the day"
		params["ResultSteady"] = false
		params["ResultContinue"] = false
		params["ResultLeaves"] = true
	}

	return params
}

// Initial Bad Morale Test
func InitialBadMorale(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Hits := params["Hits"].(float64)
	Fatigue := params["Fatigue"].(float64)
	CFatigue := params["CFatigue"].(float64)
	Leadership := params["Leadership"].(float64)
	Ejected := params["Ejected"].(bool)
	LostColors := params["LostColors"].(bool)
	ReluctantAllies := params["ReluctantAllies"].(bool)
	SQP := params["SQP"].(bool)

	params["Dice"] = ""
	params["Result"] = ""
	params["ResultHits"] = ""
	params["ResultStay"] = ""

	adder := float64(0)
	Mods, _ := list.Get(col, "InitialBadMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "CF1":
			if CFatigue == 0 {
				adder += val
			}
		case "CF2":
			if CFatigue == 1 {
				adder += val
			}
		case "CF3":
			if CFatigue == 2 {
				adder += val
			}
		case "CF4":
			if CFatigue == 3 {
				adder += val
			}
		case "ES":
			if Ejected {
				adder += val
			}
		case "FTG":
			adder += val * Fatigue
		case "HIT":
			adder += val * Hits
		case "L1":
			if Leadership == 4 {
				adder += val
			}
		case "L2":
			if Leadership == 3 {
				adder += val
			}
		case "L3":
			if Leadership == 2 {
				adder += val
			}
		case "L4":
			if Leadership == 1 {
				adder += val
			}
		case "LC":
			if LostColors {
				adder += val
			}
		case "RA":
			if ReluctantAllies {
				adder += val
			}
		case "SQP":
			if SQP {
				adder += val
			}
		}
	}

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	fid := 0
	if TotalDice >= -4 {
		fid = -4
		if TotalDice >= 0 {
			fid = 0
			if TotalDice >= 5 {
				fid = 5
				if TotalDice >= 9 {
					fid = 9
					if TotalDice >= 12 {
						fid = 12

					}
				}
			}
		}
	}
	InitialBad := list.Lookup(col, "InitialBadMorale", "Score")[fmt.Sprintf("%d", fid)]
	params["Result"] = InitialBad["Descr"]
	params["ResultHits"] = InitialBad["Hits"]
	params["ResultStay"] = InitialBad["Stay"]
	Hits += InitialBad["Hits"].(float64)
	if Hits > 10 {
		Hits = 10
	}
	params["Hits"] = Hits

	return params
}

// Check Fire Discipline in response to taking hits from small arms fire
func FireDisciplineTest(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Rating := params["Rating"].(string)
	Leader := params["Leader"].(string)
	HitsNow := params["HitsNow"].(float64)
	Hits := params["Hits"].(float64)
	BnGuns := params["BnGuns"].(bool)
	Disordered := params["Disordered"].(bool)

	params["Dice"] = ""
	params["Result"] = ""
	params["Fire"] = ""
	params["FireDisordered"] = ""
	params["PassScore"] = ""

	if HitsNow < 1 {
		params["Result"] = "No incoming fire, nothing to shoot at !"
		return params
	}

	adder := float64(0)
	Mods, _ := list.Get(col, "FireDisciplineMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "SK":
			adder += val * HitsNow
		case "HIT":
			adder += val * Hits
		case "BG":
			if BnGuns {
				adder += val
			}
		case "DIS":
			if Disordered {
				adder += val
			}
		}
	}

	switch Leader {
	case "UnInspiring":
		adder += -1
	case "Average":
		adder += 0
	case "Inspirational":
		adder += 1
	case "Charismatic":
		adder += 3
	}

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	// Get the FireDiscipline recovery table
	FireDiscipline := list.Lookup(col, "FireDisciplineTest", "Rating")[Rating]
	Pass := int(FireDiscipline["Pass"].(float64))
	Fire := int(FireDiscipline["Fire"].(float64))
	params["PassScore"] = Pass
	if TotalDice >= Pass {
		params["Result"] = "Fire Discipline Holds"
		params["ResultFire"] = false
		params["ResultFireDisordered"] = false
	} else if TotalDice >= Fire {
		params["Result"] = "Unit returns fire in good order"
		params["ResultFire"] = true
		params["ResultFireDisordered"] = false
	} else {
		params["Result"] = "Fire Discipline Fails - unit disorders and returns fire"
		params["ResultFire"] = true
		params["ResultFireDisordered"] = true
	}
	params["Disordered"] = params["ResultFireDisordered"]

	return params
}

// Check whether or not the ME accumulates fatigue this turn
func MEFatigue(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Fatigue := params["Fatigue"].(float64)
	CADefeat := params["CADefeat"].(float64)
	FF := params["FF"].(float64)
	BBM := params["BBM"].(float64)
	LostColor := params["LostColor"].(float64)
	CFatigue := params["CFatigue"].(float64)
	Leadership := params["Leadership"].(float64)
	BBOnly := params["BBOnly"].(bool)
	TookStandard := params["TookStandard"].(bool)
	NoLoss := params["NoLoss"].(bool)
	TookSP := params["TookSP"].(bool)
	TookST := params["TookST"].(bool)
	FirstBlood := params["FirstBlood"].(bool)
	ForcedMarch := params["ForcedMarch"].(bool)
	BonusImpulse := params["BonusImpulse"].(bool)
	LeaderKilled := params["LeaderKilled"].(bool)
	LeaderWounded := params["LeaderWounded"].(bool)
	CorpsCmdKilled := params["CorpsCmdKilled"].(bool)
	Mud := params["Mud"].(bool)
	Cold := params["Cold"].(bool)
	LastTurn := params["LastTurn"].(bool)

	params["Dice"] = ""
	params["Result"] = ""

	adder := float64(0)
	Mods, _ := list.Get(col, "MEFatigueMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "1ST":
			if FirstBlood {
				adder += val
			}
		case "BB":
			if BBOnly {
				adder += val
			}
		case "BI":
			if BonusImpulse {
				adder += val
			}
		case "BM":
			adder += val * BBM
		case "CD":
			adder += val * CADefeat
		case "CF":
			adder += val * CFatigue
		case "CK":
			if CorpsCmdKilled {
				adder += val
			}
		case "EC":
			if Cold {
				adder += val
			}
		case "F1":
			if TookStandard {
				adder += val
			}
		case "FF":
			adder += val * FF
		case "FM":
			if ForcedMarch {
				adder += val
			}
		case "LK":
			if LeaderKilled {
				adder += val
			}
		case "LS":
			adder += val * LostColor
		case "LW":
			if LeaderWounded {
				adder += val
			}
		case "MUD":
			if Mud {
				adder += val
			}
		case "NL":
			if NoLoss {
				adder += val
			}
		case "S1":
			if TookSP {
				adder += val
			}
		case "TS":
			if TookST {
				adder += val
			}
		}
	}

	adder += Leadership

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	params["Result"] = "No new Fatigue"
	newFatigue := 0
	if TotalDice >= 11 {
		if LastTurn {
			params["Result"] = "No new fatigue (since incurred last turn)"
		} else {
			params["Result"] = "Incurs 1 Fatigue Level (since none last turn)"
			Fatigue++
			newFatigue = 1
		}
		if TotalDice >= 15 {
			params["Result"] = "Incurs 1 Fatigue Level"
			Fatigue++
			newFatigue = 1
		}
	}

	if Fatigue > 4 {
		Fatigue = 4
	}
	params["Fatigue"] = Fatigue
	params["ResultFatigue"] = newFatigue
	params["LastTurn"] = newFatigue > 0

	return params
}

// Attempt to recover fatigue for an ME
func FatigueRecovery(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Fatigue := params["Fatigue"].(float64)
	CFatigue := params["CFatigue"].(float64)
	BBC := params["BBC"].(float64)
	Leadership := params["Leadership"].(float64)
	RestedLast := params["RestedLast"].(bool)

	params["Dice"] = ""
	params["Effect"] = ""
	params["Levels"] = ""

	adder := float64(0)
	Mods, _ := list.Get(col, "FatigueRecoveryMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "BB":
			adder += val * BBC
		case "C1":
			if CFatigue == 0 {
				adder += val
			}
		case "C2":
			if CFatigue == 1 {
				adder += val
			}
		case "C3":
			if CFatigue == 2 {
				adder += val
			}
		case "C4":
			if CFatigue == 3 {
				adder += val
			}
		case "RS":
			if RestedLast {
				adder += val
			}
		}
	}

	adder += Leadership

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	params["Effect"] = "Fatigue remains the same"
	params["RestedLast"] = false

	recovery := 0
	if TotalDice >= 8 {
		params["Effect"] = "Resting"
		params["RestedLast"] = true
		if TotalDice >= 12 {
			params["Effect"] = "Minor Fatigue Recovery"
			recovery = 1
			params["RestedLast"] = false
			if TotalDice >= 22 {
				params["Effect"] = "Major Fatigue Recovery"
				recovery = 2
			}
		}
	}

	Fatigue -= float64(recovery)
	if Fatigue < 0 {
		Fatigue = 0
	}
	params["Fatigue"] = Fatigue
	params["Levels"] = recovery

	return params
}
