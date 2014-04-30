package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/dice"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	"log"
	"strings"
)

var TroopRatings = map[string]int{
	"OldGuard":  10,
	"Guard":     9,
	"Grenadier": 8,
	"Elite":     7,
	"CrackLine": 6,
	"Veteran":   5,
	"Regular":   4,
	"Conscript": 3,
	"Landwehr":  2,
	"Militia":   1,
	"Rabble":    0,
}

func LeaderDeath(col *db.Col, params map[string]interface{}) map[string]interface{} {

	adder := float64(0)
	Charmed := int(params["Charmed"].(float64))
	Situation := int(params["Situation"].(float64))
	Nation := params["Nation"].(float64)
	Hits := params["Hits"].(float64)
	POD := params["POD"].(bool)
	Foolish := params["Foolish"].(bool)
	Rifle := params["Rifle"].(bool)
	LoseCA := params["LoseCA"].(bool)

	MaxInjury := 0
	MinInjury := 0

	switch Charmed {
	case 0: // Naturally unlucky
		adder += -2
		MinInjury = 1
	case 1: // Clumsy
		adder += -1
		MinInjury = 2
	case 3: // Debonairre
		if Foolish {
			MinInjury = 3
		}
	case 4: // Charmed
		MaxInjury = 3
		POD = false
	}

	if Charmed != 3 {
		Foolish = false
	}

	switch Nation {
	case 1, 2: // French and British, lead from the front
		if Situation >= 3 {
			adder += -1
		}
	case 3: // Cautious
		adder += 1
	}

	if Rifle {
		adder += -1
	}

	if POD {
		adder += -2
	}

	/* Types of injury

	Escape   0
	Inspired 0
	Drunk    0
	Slowed   0
	Stunned  1
	Light    2
	Serious  3
	Critical 4
	Death    5
	Captured 6

	*/
	isCA := false
	Danger := false

	switch Situation {
	case 0: // No Danger
		if Charmed == 0 {
			MaxInjury = 4 // Naturally unlucky may get a critical injury
			Danger = true
		}
	case 1: // Within cannon range of enemy
		switch Charmed {
		case 0:
			MaxInjury = 3 // the closer to danger he is, the luckier he gets !
			Danger = true
		case 1:
			MaxInjury = 2
			Danger = true
		case 2, 3, 4:
			adder += 1
			Danger = true
		}
	case 2: // Part of engaged ME
		switch Charmed {
		case 0:
			MaxInjury = 2
		case 1:
			MaxInjury = 4
		case 2, 3, 4:
			MaxInjury = 5
			adder += 1
		}
		Danger = true

	case 3: // Attached to unit which is engaged
		adder -= Hits / 2
		Danger = true
		if Foolish {
			adder += -1
		}
		if Charmed == 0 {
			MaxInjury = 4
		}
	case 4: // Attached to unit involved in close action
		adder -= Hits / 1.5
		isCA = true
		Danger = true
		if Foolish {
			adder += -2
		}
		if Charmed == 0 {
			MaxInjury = 4
		}
	case 5: // Attached to unit involved in melee
		adder -= Hits
		adder += -1
		isCA = true
		Danger = true
		if Foolish {
			adder += -3
		}
		if Charmed == 0 {
			MaxInjury = 4
		}
	case 6: // Performed a follow me
		adder -= Hits / 1.5
		adder += -2
		isCA = true
		Danger = true
		if Foolish {
			adder += -3
		}
		if Charmed == 0 {
			MaxInjury = 4
		}
	}

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	params["Result"] = ""
	params["Severity"] = ""

	if POD {
		// If Premonition of Death, then allow Death to occur
		MaxInjury = 0
	}

	if Danger {
		Injuries := list.InjuryLookup(col)
		if TotalDice < 5 {

			if MaxInjury < MinInjury {
				// Prevent infinite loop from never finding an appropriate wound
				MaxInjury = 0
			}
			for gotOne := false; !gotOne; {
				Hi := dice.DieRoll()
				Lo := dice.D6()

				params["Result"] = fmt.Sprintf("Calculating %d %d", Hi, Lo)
				params["ResultSeverity"] = "Calculating ..."
				KeyVal := uint16(Hi*10 + Lo)
				if isCA {
					KeyVal += 1000
				}
				gotOne = true
				params["Result"] = Injuries[KeyVal]["Descr"]
				params["Severity"] = Injuries[KeyVal]["Severity"]

				if MinInjury > 0 {
					// Check that the injury is of the minimum type
					switch params["Severity"] {
					case "Escape", "Inspired", "Uninspired", "Drunk", "Slowed":
						gotOne = false
					case "Stunned":
						if MinInjury > 1 {
							gotOne = false
						}
					case "Light":
						if MinInjury > 2 {
							gotOne = false
						}
					case "Serious":
						if MinInjury > 3 {
							gotOne = false
						}
					case "Critical":
						if MinInjury > 4 {
							gotOne = false
						}
					case "Death":
						if MinInjury > 5 {
							gotOne = false
						}
					}
				}
				if gotOne && MaxInjury > 0 {
					// Check that the injury is of the maximum allowed type
					switch params["Severity"] {
					case "Escape", "Inspired", "Uninspired", "Drunk", "Slowed":
						gotOne = false
					case "Stunned":
						gotOne = false
					case "Light":
						if MaxInjury < 2 {
							gotOne = false
						}
					case "Serious":
						if MaxInjury < 3 {
							gotOne = false
						}
					case "Critical":
						if MaxInjury < 4 {
							gotOne = false
						}
					case "Death":
						if MaxInjury < 5 {
							gotOne = false
						}
					case "Captured":
						if MaxInjury < 6 {
							gotOne = false
						}
					}
				}
				if !gotOne {
					log.Println("Retry", params["Severity"], "Max:", MaxInjury, "Min:", MinInjury)
				}
			}
		} else if TotalDice == 5 && LoseCA {
			params["Result"] = "The General is captured during the close action"
			params["ResultSeverity"] = "Captured"
		} else {
			params["Result"] = "The General keeps out of harms way"
		}
	} else {
		params["Result"] = "The General is in no Danger at this stage"
		params["Severity"] = ""
		params["Dice"] = ""
	}

	return params
}

// Attempt to form square under duress
func FormSquare(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Rating := params["Rating"].(string)
	Formation := params["Formation"].(string)
	Disordered := params["Disordered"].(bool)
	Attached := params["Attached"].(float64)
	Hits := params["Hits"].(float64)
	Fatigue := params["Fatigue"].(float64)
	Range := params["Range"].(float64)
	Approach := params["Approach"].(float64)
	OppCharge := params["OppCharge"].(bool)
	Action := params["Action"].(float64)

	params["PassScore"] = ""
	params["Dice"] = ""
	params["Result"] = ""
	params["ResultDisorder"] = ""

	adder := float64(0)
	Mods, _ := list.Get(col, "FormSquareMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "CA":
			if Attached == 2 {
				adder += val
			}
		case "CC":
			if Action == 1 {
				adder += val
			}
		case "DS":
			if Disordered {
				adder += val
			}
		case "FA":
			adder += val * Fatigue
		case "FL":
			if Approach == 2 {
				adder += val
			}
		case "HIT":
			adder += val * Hits
		case "LA":
			if Attached == 2 {
				adder += val
			}
		case "OC":
			if OppCharge {
				adder += val
			}
		case "RR":
			if Approach == 3 {
				adder += val
			}
		case "SG":
			if Action == 0 {
				adder += val
			}
		}
	}

	if Action == 3 {
		params["Result"] = "Unit runs for cover in disorder"
		params["ResultDisorder"] = true
		return params
	}

	switch Formation {
	case "MarchColumn":
		params["Result"] = "Unit in March Column - caught out, and disordered"
		params["ResultDisorder"] = true
		return params
	case "Skirmish":
		params["Result"] = "Unit in Skirmish Order - attempts to form Klumpen"
		params["ResultDisorder"] = true
		return params
	case "Square":
		params["Result"] = "Unit already in Square - holds position"
		return params
	}

	// Get the pass score
	gotOne := false
	PassScore := 0

	Sq, _ := list.Get(col, "FormSquare")
	for _, sq := range Sq.Data.([]interface{}) {
		mySq := sq.(map[string]interface{})

		if mySq["Rating"] == Rating && mySq["From"] == Formation {
			Field := ""
			switch Range {
			case 0:
				Field = "Grid0"
			case 1:
				Field = "Grid1"
			case 2:
				Field = "Grid1D"
			case 3:
				Field = "Grid2"
			}
			PassScore = int(mySq[Field].(float64))
			gotOne = true
			break
		}
	}
	log.Println(gotOne, PassScore)
	params["PassScore"] = PassScore

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	if TotalDice >= PassScore {
		params["Result"] = "Successfully changed formation"
		params["ResultDisorder"] = false
	} else {
		params["Result"] = "Failed to change formation, become disordered"
		params["ResultDisorder"] = true
	}

	return params
}

// Attempt to limber artillery and run off if charged
func ArtLimber(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Fatigue := params["Fatigue"].(float64)
	Enemy := params["Enemy"].(string)
	Range := params["Range"].(float64)
	Action := params["Action"].(float64)

	params["PassScore"] = ""
	params["Dice"] = ""
	params["Result"] = ""
	params["ResultDisorder"] = ""
	params["ResultEscape"] = ""
	params["ResultDistance"] = ""

	if Action == 0 {
		params["Result"] = "Stands ground to defend Battery in good order"
		params["ResultDisorder"] = false
		params["ResultEscape"] = false
		return params
	}

	Code := ""
	switch Range {
	case 0:
		Code = Enemy + "C"
	case 1:
		Code = Enemy + "S"
	case 2:
		Code = Enemy + "N"
	case 3:
		Code = Enemy + "D"
	case 4:
		params["Result"] = "Battery limbers and retires to reserve area"
		params["ResultEscape"] = true
		params["ResultDistance"] = "To Reserve"
		return params
	}
	ArtLimber := list.Lookup(col, "ArtLimber", "Code")[Code]

	adder := float64(0)
	adder -= Fatigue
	PassScore := int(ArtLimber["Score"].(float64))
	params["PassScore"] = PassScore

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	if TotalDice >= PassScore {
		params["Result"] = "Battery limbers and retires in good order"
		params["ResultEscape"] = true
		params["ResultDistance"] = ArtLimber["Flee"]
		params["ResultDisorder"] = false
	} else {
		params["Result"] = "Battery fails to limber and is caught in disorder"
		params["ResultEscape"] = false
		params["ResultDistance"] = ""
		params["ResultDisorder"] = true
	}

	return params
}

// Determine results of a Savage Street Fight
func StreetFight(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Rating := params["Rating"].(string)
	ASRating := params["ASRating"].(float64)
	Shock := params["Shock"].(bool)
	Nasty := params["Nasty"].(bool)
	Leader := params["Leader"].(string)
	ABases := params["ABases"].(float64)
	DRating := params["DRating"].(string)
	DSRating := params["DSRating"].(float64)
	DShock := params["DShock"].(bool)
	DNasty := params["DNasty"].(bool)
	DLeader := params["DLeader"].(string)
	DBases := params["DBases"].(float64)

	if ABases < 1 {
		ABases = 1
	}
	if DBases < 1 {
		DBases = 1
	}
	ARatio := ABases / DBases
	DRatio := DBases / ABases

	params["ADice"] = ""
	params["DDice"] = ""
	params["DResultHits"] = 0
	params["AResultHits"] = 0

	adder := float64(0)
	dadder := float64(0)

	switch Leader {
	case "Charismatic":
		adder += 4
	case "Inspirational":
		adder += 2
	case "Average", "Impersonal":
		adder += 1
	case "UnInspiring":
		adder -= 1
	}
	switch DLeader {
	case "Charismatic":
		dadder += 4
	case "Inspirational":
		dadder += 2
	case "Average", "Impersonal":
		dadder += 1
	case "UnInspiring":
		dadder -= 1
	}

	AR := TroopRatings[Rating]
	DR := TroopRatings[DRating]
	GradingDiff := float64((AR-DR)/2) + ASRating - DSRating

	Mods, _ := list.Get(col, "StreetMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "GR":
			adder += GradingDiff
			dadder -= GradingDiff
		case "N1":
			if ABases > DBases {
				adder += val
			}
			if DBases > ABases {
				dadder += val
			}
		case "N2":
			if ARatio >= 2 {
				adder += val
			}
			if DRatio >= 2 {
				dadder += val
			}
		case "SH":
			if Shock {
				adder += val
			}
			if DShock {
				dadder += val
			}
		case "N3":
			if ARatio >= 3 {
				adder += val
			}
			if DRatio >= 3 {
				dadder += val
			}
		case "NQ":
			if Nasty {
				adder += val
			}
			if DNasty {
				dadder += val
			}
		}
	}

	// Roll the Dice
	ADice := dice.DieRoll()
	TotalDice := ADice + int(adder)
	params["ADice"] = fmt.Sprintf("%d +%d (%d)", ADice, int(adder), TotalDice)
	DDice := dice.DieRoll()
	DTotalDice := DDice + int(dadder)
	params["DDice"] = fmt.Sprintf("%d +%d (%d)", DDice, int(dadder), DTotalDice)

	DHits := 0
	if TotalDice >= 5 {
		DHits = 1
		if TotalDice >= 9 {
			DHits = 2
			if TotalDice >= 11 {
				DHits = 3
				if TotalDice >= 13 {
					DHits = 4
					if TotalDice >= 15 {
						DHits = 5
						if TotalDice >= 17 {
							DHits = 6
							if TotalDice >= 19 {
								DHits = 7
								if TotalDice >= 22 {
									DHits = 8

								}
							}
						}
					}
				}
			}
		}
	}
	params["DResultHits"] = DHits

	AHits := 0
	if DTotalDice >= 5 {
		AHits = 1
		if DTotalDice >= 9 {
			AHits = 2
			if DTotalDice >= 11 {
				AHits = 3
				if DTotalDice >= 13 {
					AHits = 4
					if DTotalDice >= 15 {
						AHits = 5
						if DTotalDice >= 17 {
							AHits = 6
							if DTotalDice >= 19 {
								AHits = 7
								if DTotalDice >= 22 {
									AHits = 8
								}
							}
						}
					}
				}
			}
		}
	}
	params["AResultHits"] = AHits

	return params
}

// Determine results of Defensive Fire prior to close action
func DefFire(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Type := params["Type"].(string)
	Mode := params["Mode"].(float64)
	Bayonet := params["Bayonet"].(bool)
	Range := params["Range"].(float64)
	DieScore := params["DieScore"].(float64)
	Hits := params["Hits"].(float64)
	Disordered := params["Disordered"].(bool)
	Shaken := params["Shaken"].(bool)
	DGuns := params["DGuns"].(bool)

	params["Result"] = ""
	params["ResultClose"] = ""
	params["ResultDisorder"] = ""
	params["ResultFirefight"] = ""
	params["ResultFireRetire"] = ""
	params["ResultHalt"] = ""
	params["ResultRout"] = ""
	params["ResultHits"] = ""

	adder := float64(0)
	if Shaken {
		adder += 2
	}
	if Disordered {
		adder++
	}
	if Bayonet {
		adder -= 1
	}
	if DieScore+adder >= 16 {
		params["Result"] = "Excellent defensive volley halts the Attack in its tracks"
		params["ResultClose"] = ""
		if Disordered {
			params["ResultHalt"] = "1 Grid"
		} else {
			params["ResultFireRetire"] = "Halt at 1 Quad and return fire"
			params["ResultHalt"] = "1 Quad"
		}
		switch Range {
		case 0:
			params["ResultHits"] = 3
		case 1:
			params["ResultHits"] = 2
		case 2:
			params["ResultHits"] = 1
		}
		return params
	}

	// Get the lookup table
	ID := 0
	switch Type {
	case "I": // Infantry
		switch Mode {
		case 0: // Regular
			ID = 1
		case 1: // Impetus
			ID = 2
		case 2: // Storm
			ID = 3
		}
	case "S": // Shock troops
		switch Mode {
		case 0: // Regular
			ID = 4
		case 1: // Impetus
			ID = 5
		case 2: // Storm
			ID = 6
		}
	case "C": // Cavalry
		switch Mode {
		case 0: // Regular
			ID = 7
		case 1: // Impetus
			ID = 8
		}
	case "D": // DUB cavalry
		switch Mode {
		case 0: // Regular
			ID = 9
		case 1: // Impetus
			ID = 10
		}
	case "Z": // Cossacks
		ID = 11
	}
	DefFire := list.Lookup(col, "DefFire", "ID")[fmt.Sprintf("%d", ID)]
	//log.Println(DefFire)

	result := ""
	if Shaken {
		Hits *= 2
	}
	if Hits >= 1 {
		result = DefFire["Hits1"].(string)
		if Hits >= 4 {
			result = DefFire["Hits4"].(string)
			if Hits >= 6 {
				result = DefFire["Hits6"].(string)
				if Hits >= 8 {
					result = DefFire["Hits8"].(string)
					if Hits >= 10 {
						result = DefFire["Hits10"].(string)
					}
				}
			}
		}
	}
	//log.Println(result)

	// Now parse the result string
	Result := "Attacker closes with greal Elan"
	params["ResultClose"] = "Close at +1"
	Fields := strings.Fields(result)
	for _, field := range Fields {
		switch field {
		case "-1":
			params["ResultClose"] = "Close at -1"
			Result = "Attacker closes with the bayonet"
		case "-2":
			params["ResultClose"] = "Close at -2"
			Result = "Attacker eventually closes with the bayonet"
		case "-3":
			params["ResultClose"] = "Close at -3"
			Result = "Attacker struggles through to close with the bayonet"
		case "A":
			params["ResultClose"] = ""
			if dice.DieRoll() > 9 {
				if Range >= 2 {
					Range = 2
				}
				params["ResultFireRetire"] = fmt.Sprintf("Fire at %d quads, retire 2 grids", int(Range))
				params["ResultHalt"] = "2 Grids"
			} else {
				Result = "Attacker Routs"
				params["ResultRout"] = true
				switch Range {
				case 0:
					params["ResultHits"] = 2
				case 1:
					params["ResultHits"] = 1
				}
			}
		case "B":
			params["ResultClose"] = ""
			if dice.DieRoll() > 12 {
				if Range >= 2 {
					Range = 2
				}
				params["ResultFireRetire"] = fmt.Sprintf("Fire at %d quads, retire 2 grids", int(Range))
				params["ResultHalt"] = "2 Grids"
			} else {
				Result = "Attacker Routs"
				params["ResultRout"] = true
				switch Range {
				case 0:
					params["ResultHits"] = 2
				case 1:
					params["ResultHits"] = 1
				}
			}
		case "C":
			params["ResultClose"] = ""
			Result = "Cavalry attacker fires and retires"
			if Range >= 2 {
				Range = 2
			}
			params["ResultFireRetire"] = fmt.Sprintf("Fire at %d quads, retire 2 grids", int(Range))
			params["ResultHalt"] = "2 Grids"
		case "D":
			params["ResultDisorder"] = "true"
			Result = "Cavalry attacker fires and retires, disordered"
			if Disordered {
				params["Result"] = "Fallback in disorder"
				params["ResultClose"] = ""
				switch Type {
				case "I", "S":
					params["ResultHalt"] = "1 Grid"
				case "C", "D":
					params["ResultHalt"] = "2 Grids"
				case "Z":
					params["ResultHalt"] = "3 Grids"
				}
			}
		case "DC":
			params["ResultDisorder"] = "true"
			if Disordered {
				Result = "Fallback"
				params["ResultClose"] = ""
				switch Type {
				case "I", "S":
					params["ResultHalt"] = "1 Grid"
				case "C", "D":
					params["ResultHalt"] = "2 Grids"
				case "Z":
					params["ResultHalt"] = "3 Grids"
				}
			} else {
				if Range >= 2 {
					Range = 2
				}
				params["ResultFireRetire"] = fmt.Sprintf("Fire at %d quads, retire 2 grids, disordered", int(Range))
				params["ResultHalt"] = "2 Grids"
				Result = "Fire and retire"
			}
		case "E":
			if Disordered {
				Result = "Attacker Breaks"
				params["ResultClose"] = ""
				params["ResultRout"] = true
				switch Range {
				case 0:
					params["ResultHits"] = 2
				case 1:
					params["ResultHits"] = 1
				}
			}
		case "F":
			if Bayonet {
				params["ResultFireRetire"] = "Fallback 1 Quad and return fire"
				params["ResultHalt"] = "1 Quad"
				params["ResultClose"] = ""
			} else {
				params["ResultFireRetire"] = "Halt in same Quad and enter Firefight"
				params["ResultHalt"] = "Same Quad"
				params["ResultClose"] = ""
				params["ResultFirefight"] = true
			}
			Result = params["ResultFireRetire"].(string)
		case "DF":
			params["ResultDisorder"] = true
			Result += ", disordered"
			if !Bayonet || Disordered {
				params["ResultFireRetire"] = "Fallback 1 Quad and return fire, disordered"
				params["ResultHalt"] = "1 Quad"
				params["ResultClose"] = ""
				Result = "Attacker falls back 1 quad and returns fire, disordered"
			}
			if !Bayonet && !Disordered {
				params["ResultFireRetire"] = "Halt in same Quad and enter Firefight"
				params["ResultHalt"] = "Same Quad"
				params["ResultClose"] = ""
				params["ResultFirefight"] = true
			}
			Result = params["ResultFireRetire"].(string)
		case "DFG":
			params["ResultDisorder"] = true
			Result += ", disordered"
			if !Bayonet || Disordered || DGuns {
				params["ResultFireRetire"] = "Fallback 1 Quad and return fire, disordered"
				params["ResultHalt"] = "1 Quad"
				params["ResultClose"] = ""
				if DGuns {
					Result = "Cannister fire forces the Attacker back 1 quad to return fire, disordered"
				} else {
					Result = "Attacker falls back 1 quad and returns fire, disordered"
				}
			}
			if !Bayonet && !Disordered && !DGuns && !Shaken {
				params["ResultFireRetire"] = "Halt in same Quad and enter Firefight, disordered"
				Result = params["ResultFireRetire"].(string)
				params["ResultHalt"] = "Same Quad"
				params["ResultClose"] = ""
				params["ResultFirefight"] = true
			}
		case "DER":
			if Disordered {
				params["ResultDisorder"] = true
				params["ResultRout"] = true
				params["ResultClose"] = ""
				Result = "Attacker Breaks in Disgrace, Shame and Disorder"
				switch Range {
				case 0:
					params["ResultHits"] = 2
				case 1:
					params["ResultHits"] = 1
				}
			} else {
				params["ResultDisorder"] = true
				params["ResultHalt"] = "1 Grid"
				params["ResultClose"] = ""
				Result = "Attacker Retreats 1 Grid to hold their position, disordered"
			}
		case "G":
			if DGuns {
				params["ResultDisorder"] = true
				params["ResultFireRetire"] = "Fallback 1 Quad and return fire, disordered"
				params["ResultHalt"] = "1 Quad"
				params["ResultClose"] = ""
				Result = "Cannister fire forces the Attacker back 1 quad to return fire, disordered"
			}
		case "DG":
			params["ResultDisorder"] = true
			Result += ", disordered"
			if Disordered || DGuns {
				params["ResultFireRetire"] = "Fallback 1 Quad and return fire, disordered"
				params["ResultHalt"] = "1 Quad"
				params["ResultClose"] = ""
				if DGuns {
					Result = "Cannister fire forces the Attacker back 1 quad to return fire, disordered"
				} else {
					Result = "Attacker is forced back 1 quad to return fire, disordered"
				}
			}
		case "FG":
			if Bayonet && params["ResultClose"] != "" {
				Result = "Attacker closes with the Bayonet"
			} else {
				params["ResultFireRetire"] = "Fallback 1 Quad and return fire"
				params["ResultHalt"] = "1 Quad"
				params["ResultClose"] = ""
				Result = "Attacker falls back and returns fire"
				if !Disordered && !DGuns && !Shaken {
					params["ResultFireRetire"] = "Halt in same Quad and enter Firefight"
					Result = params["ResultFireRetire"].(string)
					params["ResultHalt"] = "Same Quad"
					params["ResultClose"] = ""
					params["ResultFirefight"] = true
				}

			}
		case "DEFG":
			params["ResultDisorder"] = true
			Result += ", disordered"
			if !Disordered {
				if Bayonet && params["ResultClose"] != "" && !DGuns {
					Result = "Attacker closes with the Bayonet, disordered"
				} else {
					params["ResultFireRetire"] = "Fallback 1 Quad and return fire disordered"
					params["ResultHalt"] = "1 Quad"
					params["ResultClose"] = ""
					Result = "Attacker falls back and returns fire disordered"
				}
			} else {
				Result = "Attacker Breaks"
				params["ResultClose"] = ""
				params["ResultRout"] = true
				switch Range {
				case 0:
					params["ResultHits"] = 2
				case 1:
					params["ResultHits"] = 1
				}
			}
		case "H":
			if !Bayonet {
				params["ResultHalt"] = "Same Quad"
				params["ResultFirefight"] = true
				params["ResultClose"] = ""
				Result = "Halt at close range and enter firefight"
			}
		case "R":
			params["ResultHalt"] = "1 Grid"
			params["ResultClose"] = ""
			Result = "Attacker Retreats 1 Grid in good order"
		case "Z":
			params["ResultHalt"] = "2 Grid"
			params["ResultClose"] = ""
			Result = "Attacker Retreats 2 Grids in good order"
		case "DR":
			params["ResultHalt"] = "1 Grid"
			params["ResultClose"] = ""
			params["ResultDisorder"] = true
			Result = "Attacker Retreats 1 Grid in disorder"
		case "DZ":
			params["ResultHalt"] = "2 Grid"
			params["ResultClose"] = ""
			params["ResultDisorder"] = true
			Result = "Attacker Retreats 2 Grids in disorder"
		case "X":
			Result = "Attacker Breaks"
			params["ResultClose"] = ""
			params["ResultRout"] = true
			switch Range {
			case 0:
				params["ResultHits"] = 3
			case 1:
				params["ResultHits"] = 2
			case 2:
				params["ResultHits"] = 1
			}
		}
	}

	if Shaken && params["ResultClose"] != "" {
		Result = "Shaken Attacker halts at 1 Quad and fires in disorder"
		params["ResultClose"] = ""
		params["ResultHalt"] = "1 Quad"
		params["ResultDisorder"] = true
		params["ResultFireRetire"] = "Halts at 1 Quad and fires, disordered"
	}
	params["Result"] = Result

	return params
}
