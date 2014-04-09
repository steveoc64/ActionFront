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

	retval := make(map[string]interface{})

	retval["DepRating"] = params["DepRating"]
	retval["DepState"] = params["DepState"]
	retval["MarchCol"] = params["MarchCol"]
	retval["Darkness"] = params["Darkness"]
	retval["Choke"] = params["Choke"]
	retval["Mud"] = params["Mud"]
	retval["Fog"] = params["Fog"]
	retval["Grids"] = params["Grids"]
	retval["Dice"] = 0
	retval["DieMods"] = 0
	retval["Result"] = ""

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

	retval["DieMods"] = adjust
	d := dice.DieRoll()
	Score := d + adjust

	retval["Dice"] = fmt.Sprintf("%d + %d = %d", d, adjust, Score)

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
		retval["Result"] = "No Change"
	} else {
		resString := params["DepState"].(string) + " -> "
		switch depState {
		case 1:
			retval["DepState"] = "Deployed"
		case 2:
			retval["DepState"] = "Bde Out"
		case 3:
			retval["DepState"] = "Deploying"
		case 4:
			retval["DepState"] = "Condensed Col"
		case 5:
			retval["DepState"] = "Regular Col"
		case 6:
			retval["DepState"] = "Extended Col"
		}
		retval["Result"] = resString + retval["DepState"].(string)
	}

	return retval
}

// For a given set of parameters, calculate the Tactical Move stats, and return this as a result set
func TacMove(col *db.Col, params map[string]interface{}) map[string]interface{} {

	retval := make(map[string]interface{})

	retval["UnitType"] = params["UnitType"]
	retval["DrillBook"] = params["DrillBook"]
	retval["Formation"] = params["Formation"]
	retval["FormationTo"] = params["FormationTo"]
	retval["Terrain"] = params["Terrain"]
	retval["Extra"] = params["Extra"]
	retval["Mud"] = params["Mud"]
	retval["Marsh"] = params["Marsh"]
	retval["Diagonal"] = params["Diagonal"]
	retval["LoWall"] = params["LoWall"]
	retval["HiWall"] = params["HiWall"]
	retval["LtWood"] = params["LtWood"]
	retval["HvWood"] = params["HvWood"]
	retval["Weather"] = params["Weather"]
	retval["Accumulated"] = params["Accumulated"]
	retval["Trained"] = params["Trained"]
	retval["Disorder"] = false
	retval["Fire"] = false
	retval["Quadrants"] = 0
	retval["Inches"] = 0
	retval["Frontage"] = 0
	retval["SK"] = ""

	// Get the TacMove record for this unit type
	TacMoves, _ := list.Get(col, "TacMove")
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
	for _, tacMove := range TacMoves.Data.([]interface{}) {
		myTacMove := tacMove.(map[string]interface{})

		// Adjust for Type of unit
		if myTacMove["UnitType"] == params["UnitType"] {
			baseMove = myTacMove["Move"].(float64)
		}
	}

	// Adjust for the drill type and formation, applies to infantry only
	if isInf {
		for _, drill := range Drills.Data.([]interface{}) {
			myDrill := drill.(map[string]interface{})
			if myDrill["Name"] == params["DrillBook"] {

				// Get the Drill entry for the start formation
				entries := myDrill["Entries"].(map[string]interface{})
				startDrill := entries[params["Formation"].(string)].(map[string]interface{})

				frontage = startDrill["FR"].(float64)
				retval["SK"] = fmt.Sprintf("%d / %d", int(startDrill["OO"].(float64)), int(startDrill["SK"].(float64)))
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
						retval["FormationTo"] = retval["Formation"]
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
	retval["Inches"] = math.Trunc(inches)
	inchesPerQuad := float64(5)
	if params["Diagonal"].(bool) {
		inchesPerQuad = 7
	}

	totalInches := inches + startInches
	quads := totalInches / inchesPerQuad
	if isSK {
		disorder = false
	}
	retval["Quadrants"] = math.Trunc(quads)
	retval["Accumulated"] = math.Trunc(math.Mod(totalInches, inchesPerQuad))
	retval["Disorder"] = disorder
	retval["Frontage"] = frontage
	retval["Fire"] = canFire

	return retval
}
