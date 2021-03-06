package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/dice"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	"log"
	"math"
)

// For a given set of parameters, calculate the GTMove, and return this as a result set
func GTMove(col *db.Col, params map[string]interface{}) map[string]interface{} {

	var baseMove float64

	// Try this loop using the Lookup service

	GTMoveLookup := list.Lookup(col, "GTMove", "METype")
	GTMove := GTMoveLookup[params["METype"].(string)]

	// We now have the correct GT Move record
	switch params["DeploymentState"] {
	case "Deployed":
		baseMove = GTMove["D1"].(float64)
		if params["MarchOrder"].(bool) {
			baseMove += 4
		}
	case "Bde Out":
		baseMove = GTMove["D2"].(float64)
	case "Deploying":
		baseMove = GTMove["D3"].(float64)
	case "Condensed Col":
		baseMove = GTMove["D4"].(float64)
	case "Regular Col":
		baseMove = GTMove["D5"].(float64)
	case "Extended Col":
		baseMove = GTMove["D6"].(float64)
	}

	// Lets see if we have a forced march on our hands
	if params["Forced"].(bool) {
		checkf := GTMoveLookup["Forced March"]
		switch params["DeploymentState"] {
		case "Deployed":
			baseMove += checkf["D1"].(float64)
		case "Bde Out":
			baseMove += checkf["D2"].(float64)
		case "Deploying":
			baseMove += checkf["D3"].(float64)
		case "Condensed Col":
			baseMove += checkf["D4"].(float64)
		case "Regular Col":
			baseMove += checkf["D5"].(float64)
		case "Extended Col":
			baseMove += checkf["D6"].(float64)
		}
	}

	acc := params["Accumulated"].(float64)
	turns := 1.0

	// Get the appropriate weather modifier
	WeatherLookup := list.Lookup(col, "Weather", "Code")
	Weather := WeatherLookup[params["Weather"].(string)]
	if Weather["Code"] == params["Weather"] {
		// We now have the appropriate weather as well

		baseMove = baseMove * Weather["Move"].(float64) / 10.0
	}

	baseMove *= turns
	inchesPerGrid := 10.0
	if params["Diagonal"].(bool) {
		inchesPerGrid = 15.0
	}
	params["Inches"] = math.Trunc(baseMove)
	params["Distance"] = math.Trunc((baseMove + acc) / inchesPerGrid)
	params["Accumulated"] = math.Trunc(math.Mod(baseMove+acc, inchesPerGrid))

	return params
}

// For a given set of parameters, calculate the Deployment stats, and return this as a result set
func Deployment(col *db.Col, params map[string]interface{}) map[string]interface{} {

	params["Dice"] = 0
	params["DieMods"] = 0
	params["Result"] = ""

	// get the adjustment bonus for this rating

	adjust := 0

	DepMods, _ := list.Get(col, "DeploymentMod")

	for _, depMod := range DepMods.Data.([]interface{}) {
		myDepMod := depMod.(map[string]interface{})

		val := int(myDepMod["Value"].(float64))

		// Adjust for Type of unit
		if myDepMod["Descr"] == params["DepRating"] {
			adjust += val
		}

		// Adjust for other known conditions
		if params["Mud"].(bool) && myDepMod["Code"] == "MUD" {
			adjust += val
		}
		if params["Fog"].(bool) && myDepMod["Code"] == "FOG" {
			adjust += val
		}
		if params["Choke"].(bool) && myDepMod["Code"] == "CP" {
			adjust += val
		}
		if params["Darkness"].(bool) && myDepMod["Code"] == "DK" {
			adjust += val
		}
		if params["Grids"].(float64) != 0 && myDepMod["Code"] == "MV" {
			adjust += (int(params["Grids"].(float64)) * val)
		}
	}

	params["DieMods"] = adjust
	d := dice.DieRoll()
	Score := d + adjust

	params["Dice"] = fmt.Sprintf("%d + %d = %d", d, adjust, Score)

	// Convert the DepState to a number
	depState := 0
	switch params["DepState"] {
	case "Deployed":
		depState = 1
	case "Bde Out":
		depState = 2
	case "Deploying":
		depState = 3
	case "Condensed Col":
		depState = 4
	case "Regular Col":
		depState = 5
	case "Extended Col":
		depState = 6
	}

	direction := -1
	if params["MarchCol"].(bool) {
		direction = 1
	}
	change := 0

	// Compare the adjusted die roll to the score needed
	if Score >= 1 {
		change = 1
		if Score >= 10 {
			change = 2
			if Score >= 16 {
				change = 3
			}
		}
	}
	change *= direction
	depState += change
	if depState < 1 {
		depState = 1
	}
	if depState > 6 {
		depState = 6
	}
	if change == 0 {
		params["Result"] = "No Change"
	} else {
		resString := params["DepState"].(string) + " -> "
		switch depState {
		case 1:
			params["DepState"] = "Deployed"
		case 2:
			params["DepState"] = "Bde Out"
		case 3:
			params["DepState"] = "Deploying"
		case 4:
			params["DepState"] = "Condensed Col"
		case 5:
			params["DepState"] = "Regular Col"
		case 6:
			params["DepState"] = "Extended Col"
		}
		params["Result"] = resString + params["DepState"].(string)
	}

	return params
}

// For a given set of parameters, calculate the Tactical Move stats, and return this as a result set
func TacMove(col *db.Col, params map[string]interface{}) map[string]interface{} {

	params["Disorder"] = false
	params["Fire"] = false
	params["Quadrants"] = 0
	params["Inches"] = 0
	params["Frontage"] = 0
	params["SK"] = ""

	// Get the TacMove record for this unit type
	TacMoves := list.Lookup(col, "TacMove", "UnitType")
	Drills, _ := list.Get(col, "Drill")
	FormationChanges, _ := list.Get(col, "FormationChange")
	baseMove := float64(10)
	multiplier := float64(1)
	adder := float64(0)
	frontage := float64(1)
	isInf := false
	isCav := false
	isArt := false
	isSK := false
	disorder := params["Disorder"].(bool)
	canFire := true

	if params["Formation"].(string) == "Skirmish" {
		isSK = true
	}

	switch params["UnitType"].(string) {
	case "Artillery":
		isArt = true
	case "Cavalry", "LightCav":
		isCav = true
	default:
		isInf = true
	}

	// Adjust for the unit type
	baseMove = TacMoves[params["UnitType"].(string)]["Move"].(float64)

	// Adjust for the drill type and formation, applies to infantry only
	if isInf {
		for _, drill := range Drills.Data.([]interface{}) {
			myDrill := drill.(map[string]interface{})
			if myDrill["Name"] == params["DrillBook"] {

				// Get the Drill entry for the start formation
				entries := myDrill["Entries"].(map[string]interface{})
				startDrill := entries[params["Formation"].(string)].(map[string]interface{})

				frontage = startDrill["FR"].(float64)
				params["SK"] = fmt.Sprintf("%d / %d", int(startDrill["OO"].(float64)), int(startDrill["SK"].(float64)))
				multiplier = startDrill["EF"].(float64) / 10
			}
		}

		// If changing formation - do some major adjustments
		ffrom := params["Formation"].(string)
		fto := params["FormationTo"].(string)
		if ffrom == "Line" && fto != "Line Left" ||
			ffrom != "Line" && ffrom != fto {

			for _, fchange := range FormationChanges.Data.([]interface{}) {
				myfc := fchange.(map[string]interface{})
				era := "DIV"
				if params["DrillBook"].(string) == "Old School" {
					era = "AR"
				}
				if myfc["Era"].(string) == era &&
					myfc["From"].(string) == params["Formation"].(string) &&
					myfc["To"].(string) == params["FormationTo"].(string) {
					log.Println("Change record that applies", myfc)

					fceffect := float64(0)
					switch params["Trained"].(string) {
					case "UnTrained":
						fceffect = myfc["Untrained"].(float64)
					default:
						fceffect = myfc["Trained"].(float64)
					}

					switch fceffect {
					case -3:
						multiplier = 0
						canFire = false
						disorder = true
					case -2:
						multiplier = 0
						canFire = false
					case -1:
						multiplier = 0
						canFire = true
					case 0:
						// Not allowed
						params["FormationTo"] = params["Formation"]
					default:
						baseMove = fceffect
					}
				}
			}

		}
	}

	// Adjust for the terrain
	switch params["Terrain"].(string) {
	case "Marchfeld":
		multiplier *= 1.1
	case "Rolling":
		multiplier *= 1.0
	case "Rough":
		multiplier *= 0.8
		frontage -= 1
	case "Hill":
		if isCav {
			multiplier *= 0.3
			frontage -= 1
			canFire = false
		} else if isArt {
			multiplier *= 0.2
			frontage = 1
			canFire = false
		} else {
			multiplier *= 0.4
			frontage -= 2
			disorder = true
		}
	case "Town":
		frontage = 1
		if isCav {
			multiplier *= 0.2
			canFire = false
		} else if isArt {
			multiplier *= 0.1
			canFire = false
		} else {
			multiplier *= 0.5
			disorder = true
		}
	}

	// Adjust for terrain effects
	if params["LtWood"].(bool) {
		if isArt {
			multiplier *= 0.5
		} else if isCav {
			multiplier *= 0.5
			disorder = true
		} else {
			multiplier *= 0.7
			if !disorder {
				disorder = dice.Percent(40)
			}
			frontage -= 2
			if canFire {
				canFire = dice.Percent(60)
			}
		}
	}
	if params["HvWood"].(bool) {
		if isArt {
			multiplier *= 0
		} else if isCav {
			multiplier *= 0
		} else {
			multiplier *= 0.5
			if !disorder {
				disorder = dice.Percent(80)
			}
			frontage = 1
			if canFire {
				canFire = dice.Percent(30)
			}
		}
	}
	if params["Mud"].(bool) {
		if isArt {
			multiplier *= 0.2
			canFire = false
		} else if isCav {
			multiplier *= 0.3
			if !disorder {
				disorder = dice.Percent(50)
			}
			canFire = false
		} else {
			multiplier *= 0.5
			if !disorder {
				disorder = dice.Percent(30)
			}
			frontage -= 1
			if canFire {
				canFire = dice.Percent(60)
			}
		}
	}
	if params["Marsh"].(bool) {
		if isArt {
			multiplier *= 0
		} else if isCav {
			multiplier *= 0.6
			if !disorder {
				disorder = dice.Percent(50)
			}
			canFire = false
		} else {
			multiplier *= 0.7
			if !disorder {
				disorder = dice.Percent(50)
			}
			frontage = 2
			if canFire {
				canFire = dice.Percent(80)
			}
		}
	}
	if params["LoWall"].(bool) {
		if isArt {
			multiplier *= 0
		} else if isCav {
			adder -= 5
			canFire = false
		} else {
			adder -= 1
			if !disorder {
				disorder = dice.Percent(30)
			}
			if canFire {
				canFire = dice.Percent(90)
			}
		}
	}
	if params["HiWall"].(bool) {
		if isArt {
			multiplier *= 0
		} else if isCav {
			multiplier *= 0
		} else {
			adder -= 3
			if !disorder {
				disorder = dice.Percent(60)
			}
			if canFire {
				canFire = dice.Percent(60)
			}
		}
	}

	// If infantry, and not in Skirmish order, apply march pace
	if isInf && !isSK {
		switch params["Extra"].(float64) {
		case 0:
			if disorder {
				// small chance of recovering any disorder if advancing at a slow march
				disorder = dice.Percent(70)
			}
		case 1:
			adder += 1
			// small chance of becoming disordered, depending on terrain
			switch params["Terrain"].(string) {
			case "Marchfeld":
				if !disorder {
					disorder = dice.Percent(10)
				}
			case "Rolling":
				if !disorder {
					disorder = dice.Percent(20)
				}
			case "Rough":
				if !disorder {
					disorder = dice.Percent(30)
				}
			}

		case 3:
			adder += 3
			// Good chance of becoming disordered, depending on terrain
			switch params["Terrain"].(string) {
			case "Marchfeld":
				if !disorder {
					disorder = dice.Percent(30)
				}
			case "Rolling":
				if !disorder {
					disorder = dice.Percent(60)
				}
			case "Rough":
				if !disorder {
					disorder = dice.Percent(90)
				}
			}

		}
	}

	if isCav {
		switch params["Extra"].(float64) {
		case 0:
			adder -= 4
			// slow trotting pace, to ensure good order
			if disorder {
				disorder = dice.Percent(60)
			}
		case 1:
			if !disorder {
				disorder = dice.Percent(20)
			}
		case 3:
			adder += 6
			if !disorder {
				disorder = dice.Percent(40)
			}
		}
	}

	// Perform final calculations
	startInches := params["Accumulated"].(float64)
	inches := (baseMove * multiplier) + adder
	params["Inches"] = math.Trunc(inches)
	inchesPerQuad := float64(5)
	if params["Diagonal"].(bool) {
		inchesPerQuad = 7
	}

	totalInches := inches + startInches
	quads := totalInches / inchesPerQuad
	if isSK {
		disorder = false
	}
	params["Quadrants"] = math.Trunc(quads)
	params["Accumulated"] = math.Trunc(math.Mod(totalInches, inchesPerQuad))
	params["Disorder"] = disorder
	params["Frontage"] = frontage
	params["Fire"] = canFire

	return params
}

// For a given set of parameters, calculate the Artillery Move stats, and return this as a result set
func ArtyMove(col *db.Col, params map[string]interface{}) map[string]interface{} {

	MoveType := params["MoveType"].(string)
	MoveWeight := params["MoveWeight"].(string)
	Accumulated := params["Accumulated"].(float64)
	Pace := params["Pace"].(float64)
	Terrain := params["Terrain"].(float64)
	params["Quadrants"] = 0
	params["Inches"] = 0
	params["HorseLoss"] = 0

	// Get the right movement record
	baseMove := int(0)
	ArtyMoves, _ := list.Get(col, "ArtyMove")
	for _, move := range ArtyMoves.Data.([]interface{}) {
		myMove := move.(map[string]interface{})
		if myMove["Class"] == MoveType && myMove["Weight"] == MoveWeight {
			log.Println(myMove)
			switch Pace {
			case 1:
				baseMove = int(myMove["Prolong"].(float64))
			case 2:
				baseMove = int(myMove["Regular"].(float64))
			case 3:
				baseMove = int(myMove["Gallop"].(float64))

				// Calc horse loss due to Galloping
				HorseLosses := list.Lookup(col, "ArtyHorseLoss", "Terrain")
				useTerrain := ""
				switch Terrain {
				case 1:
					useTerrain = "Marchfeld"
				case 2:
					useTerrain = "Rolling"
				case 3:
					useTerrain = "Rough"
				}
				HorseLossChance := int(HorseLosses[useTerrain]["Loss"].(float64))
				if dice.DieRoll() <= HorseLossChance {
					params["HorseLoss"] = 1
				}
			}
		}
	}

	if baseMove == 0 {
		log.Println("Unknown class/weight", MoveType, MoveWeight)
		return params
	}

	// Perform final calculations
	params["Inches"] = baseMove
	inchesPerQuad := float64(5)
	if params["Diagonal"].(bool) {
		inchesPerQuad = 7
	}

	totalInches := float64(baseMove + int(Accumulated))
	quads := totalInches / inchesPerQuad
	params["Quadrants"] = math.Trunc(quads)
	params["Accumulated"] = math.Trunc(math.Mod(totalInches, inchesPerQuad))

	return params
}

// Withdraw or Relocate an artillery unit
func ArtyRelocate(col *db.Col, params map[string]interface{}) map[string]interface{} {

	GunneryClass := ""
	switch params["GunneryClass"].(float64) {
	case 0:
		GunneryClass = "Guard"
	case 1:
		GunneryClass = "Class I"
	case 2:
		GunneryClass = "Class II"
	case 3:
		GunneryClass = "Class III"
	}

	Action := params["Action"].(float64)
	Horses := params["Horses"].(float64)
	Fatigue := params["Fatigue"].(float64)
	Attached := params["Attached"].(float64)
	Attempt := params["Attempt"].(float64)
	params["Dice"] = ""
	params["Result"] = ""

	adder := float64(0)
	RMods, _ := list.Get(col, "ArtyRelocateMod")
	for _, rmod := range RMods.Data.([]interface{}) {
		myRMod := rmod.(map[string]interface{})

		code := myRMod["Code"].(string)
		val := myRMod["Value"].(float64)
		switch code {
		case "MD":
			if params["Mud"].(bool) {
				adder += val
			}
		case "FT":
			adder += val * Fatigue
		case "AT":
			adder += val * (Attempt - 1)
		case "LA":
			if Attached == 1 {
				adder += val
			}
		case "CA":
			if Attached == 2 {
				adder += val
			}
		case "AA":
			if Attached == 3 {
				adder += val
			}
		}
	}

	// Now get the relocation record
	Relocations := list.Lookup(col, "ArtyRelocate", "Class")
	Relocation := Relocations[GunneryClass]
	Field := ""
	ActionString := ""
	if Fatigue >= 4 {
		Action = 0
		params["Action"] = 0
	}
	switch Action {
	case 0: // Withdraw
		Field = "W"
		ActionString = "Withdraw to Reserve"
	case 1: // Relocate
		Field = "R"
		ActionString = "Relocate"
	}
	FieldName := fmt.Sprintf("%s%d", Field, int(Horses))
	log.Println("Field Name", FieldName)
	Value := int(Relocation[FieldName].(float64))
	log.Println("Value Needed", Value)
	params["ScoreNeeded"] = Value
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)
	if TotalDice >= Value {
		params["Result"] = fmt.Sprintf("Battery will %s", ActionString)
		params["Attempt"] = 1
		if Action == 0 {
			Fatigue++
			params["Fatigue"] = Fatigue
		}
	} else {
		params["Result"] = fmt.Sprintf("Failed to %s", ActionString)
		Attempt++
		if Attempt > 3 {
			Attempt = 3
		}
		params["Attempt"] = Attempt
	}

	return params
}

// Double Team a Battery
func DoubleTeam(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Type := params["Type"].(float64)
	adder := float64(0)
	Mods, _ := list.Get(col, "DoubleTeamMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "French Guard":
			if Type == 1 {
				adder += val
			}
		case "Horse Arty":
			if Type == 2 {
				adder += val
			}
		case "French":
			if Type == 3 {
				adder += val
			}
		case "British":
			if Type == 4 {
				adder += val
			}
		case "Other":
			if Type == 5 {
				adder += val
			}
		}
	}
	params["ScoreNeeded"] = 11
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)
	if TotalDice >= 11 {
		params["Result"] = "Successfully Double Teamed the Battery"
		params["Exhausted"] = false
	} else {
		params["Result"] = "Failed to Double Team the Battery"
		params["Exhausted"] = true
	}

	return params
}

// Attempt to Recover Guns
func RecoverGuns(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Owner := params["Owner"].(float64)
	adder := float64(0)
	Mods, _ := list.Get(col, "ArtFateMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "NE":
			if params["NE"].(bool) {
				adder += val
			}
		case "CA":
			if params["CA"].(bool) {
				adder += val
			}
		case "EN":
			if params["EN"].(bool) {
				adder += val
			}
		}
	}
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	Fates := list.Lookup(col, "ArtFate", "Situation")
	params["ResultDisabled"] = false
	params["ResultRecovered"] = false

	switch Owner {
	case 1: // Own guns - try to recover
		params["ScoreNeeded"] = Fates["Friendly"]["Score"]
		if TotalDice >= int(Fates["Friendly"]["Score"].(float64)) {
			params["Result"] = "Successfully Recovered Guns"
			params["ResultRecovered"] = true
		} else {
			params["Result"] = "Failed to Recover Guns"
		}
	case 2: // Enemy Guns
		params["ScoreNeeded"] = Fates["Capture Enemy"]["Score"]
		if TotalDice >= int(Fates["Capture Enemy"]["Score"].(float64)) {
			params["Result"] = "Successfully Captured Enemy Guns"
			params["ResultRecovered"] = true
		} else if TotalDice >= int(Fates["Disable Enemy"]["Score"].(float64)) {
			params["Result"] = "Disabled Enemy Guns"
			params["ResultDisabled"] = true
		} else {
			params["Result"] = "Failed to deal with Enemy Guns"
		}

	}

	return params
}

// Attempt to Move Skirmishers away from parent unit
func SKRelocate(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Ammo := params["Ammo"].(float64)
	Hits := params["Hits"].(float64)
	Fatigue := params["Fatigue"].(float64)
	Range := int(params["Range"].(float64))
	Terrain := params["Terrain"].(float64)
	Leader := params["Leader"].(string)
	Rating := params["Rating"].(string)

	// Set default results
	params["Result"] = ""
	params["ResultNoMove"] = false
	params["ResultRetire"] = false
	params["ResultMove"] = false
	params["ResultBold"] = false

	// Get the relocation record
	SKRelocate := list.Lookup(col, "SKRelocate", "Rating")[Rating]
	SKSupport := list.Lookup(col, "SKSupport", "Mode")

	// Apply all the modifiers
	adder := float64(0)
	Mods, _ := list.Get(col, "SKRelocateMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "AM":
			if Ammo >= 1 {
				adder += val
			}
		case "HIT":
			adder += val * Hits
		case "FT":
			adder += val * Fatigue
		case "UL":
			if Leader == "UnInspiring" {
				adder += val
			}
		case "AL":
			if Leader == "Average" {
				adder += val
			}
		case "IL":
			if Leader == "Inspirational" {
				adder += val
			}
		case "CL":
			if Leader == "Charismatic" {
				adder += val
			}
		case "BD":
			if params["Bold"].(bool) {
				adder += val
			}

		}
	}

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)
	params["Bold"] = false

	if TotalDice < int(SKRelocate["Retire"].(float64)) {
		params["Result"] = "No Move, Lose Initiative"
		params["ResultNoMove"] = true
	} else if TotalDice < int(SKRelocate["Move"].(float64)) {
		supportDistance := SKSupport["Normal"]
		distance := 0
		switch Terrain {
		case 0:
			distance = int(supportDistance["Marchfeld"].(float64))
		case 1:
			distance = int(supportDistance["Rolling"].(float64))
		case 2:
			distance = int(supportDistance["Rough"].(float64))
		}

		if distance > Range {
			params["Result"] = fmt.Sprintf("Must fallback %d quadrants, to within %d of parent unit", distance-Range, distance)
			params["Range"] = distance
		}

		params["Result"] = "May hold position, or retire"
		params["ResultRetire"] = true
	} else if TotalDice < int(SKRelocate["Bold"].(float64)) {
		// Can move - work out how far
		supportDistance := SKSupport["Normal"]
		distance := 0
		switch Terrain {
		case 0:
			distance = int(supportDistance["Marchfeld"].(float64))
		case 1:
			distance = int(supportDistance["Rolling"].(float64))
		case 2:
			distance = int(supportDistance["Rough"].(float64))
		}

		params["ResultRetire"] = true
		if distance > Range {
			params["Result"] = fmt.Sprintf("Advance %d quadrants, keeping within %d of parent unit", distance-Range, distance)
			params["ResultMove"] = true
		} else if distance == Range {
			params["Result"] = fmt.Sprintf("At their limit of %d quadrants from parent unit", distance)
			params["ResultMove"] = false
		} else {
			params["Result"] = fmt.Sprintf("Fallback to the limit of %d quadrants from parent unit", distance)
			params["ResultMove"] = true
		}
		params["Range"] = distance
	} else {
		// Can move - work out how far
		supportDistance := SKSupport["Bold"]
		distance := 0
		switch Terrain {
		case 0:
			distance = int(supportDistance["Marchfeld"].(float64))
		case 1:
			distance = int(supportDistance["Rolling"].(float64))
		case 2:
			distance = int(supportDistance["Rough"].(float64))
		}

		params["ResultRetire"] = true
		if distance > Range {
			params["Result"] = fmt.Sprintf("May Boldly Advance %d quadrants, keeping within %d of parent unit", distance-Range, distance)
			params["ResultMove"] = true
		} else if distance == Range {
			params["Result"] = fmt.Sprintf("At their Bold limit of %d quadrants from parent unit", distance)
			params["ResultMove"] = false
		} else {
			params["Result"] = fmt.Sprintf("Fallback to the Bold limit of %d quadrants from parent unit", distance)
			params["ResultMove"] = true
		}
		params["ResultBold"] = true
		params["Bold"] = true
		params["Range"] = distance
	}

	return params
}

// Attempt Occupy or exit a BUA
func BUAMove(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Action := params["Action"].(string)
	SRating := params["SRating"].(string)
	Hits := params["Hits"].(float64)
	Fatigue := params["Fatigue"].(float64)
	Special := params["Special"].(string)
	CA := params["CA"].(string)
	UnitsMoved := params["UnitsMoved"].(float64)
	Rain := params["Rain"].(bool)
	Cold := params["Cold"].(bool)

	// Set default results
	params["Result"] = ""
	params["ResultOrdered"] = false

	// Get the lookup records
	BUAMove := list.Lookup(col, "BUAMove", "Rating")[SRating]

	// Apply all the modifiers
	adder := float64(0)
	Mods, _ := list.Get(col, "BUAMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "UN":
			adder += val * UnitsMoved
		case "HIT":
			adder += val * Hits
		case "FT":
			adder += val * Fatigue
		case "LA":
			if CA == "LA" {
				adder += val
			}
		case "CA":
			if CA == "CA" {
				adder += val
			}
		case "COLD":
			if Cold {
				adder += val
			}
		case "RAIN":
			if Rain {
				adder += val
			}
		case "RU":
			if Special == "RU" {
				adder += val
			}
		case "AU":
			if Special == "AU" {
				adder += val
			}
		case "JN":
			if Special == "JN" {
				adder += val
			}
		}
	}

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	Ordered := int(BUAMove["Ordered"].(float64))
	Exit := int(BUAMove["Exit"].(float64))

	switch Action {
	case "O":
		if TotalDice < Ordered {
			params["Result"] = "Occupy structure in a disordered state"
			params["ResultOrdered"] = false
		} else {
			params["Result"] = "Occupy structure in good order"
			params["ResultOrdered"] = true
		}
	case "X":
		if TotalDice < Exit {
			params["Result"] = "Remain in the structure in a disordered state"
			params["ResultOrdered"] = false
		} else if TotalDice < Ordered {
			params["Result"] = "Exit the structure in a disordered state"
			params["ResultOrdered"] = false
		} else {
			params["Result"] = "Exit the structure in good order"
			params["ResultOrdered"] = true
		}

	}

	return params
}
